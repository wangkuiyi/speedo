SpeedoVision uses Caffe to train CNNs that maps what a car sees to control commands.

To accomplish this goal, there are several steps. 

1. Generate "lane alike" pictures with backgrounds and lanes that have some different 
angels with the horizon. Train a CNN to see whether it can learn the angel label from 
the images. 

2. If Step 1 performs well, we can extend the training using real photos from the 
environment. We will train CNNs and RNNs to map the image to control sequences. 
This step is similar to [what Princeton is doing](deepdriving.cs.princeton.edu).

3. Deploy a good model on a cheap chip that can real time scoring the photos in 24Hz~36Hz.
We may need to optimize it, otherwise the chip is not powerful enough. 

### Image generation for Step 1

image_gen.pde is a [processing](www.processing.org) program. It generates images with 
different background colors and lane angels for the simulation training purposes. 

### Caffe quick start guide

There are several dependencies must be installed before installing Caffe. But it's straightforward to
follow the installation guide [here](http://caffe.berkeleyvision.org/installation.html#compilation).
If trying python interface, Python 2.7 is recommended rather than 3.5 because it's the default python version in Caffe. 

On MacOS, GPU mode may not work well even if you install the CUDA package and compile Caffe as instructed. 
[Here](https://github.com/BVLC/caffe/issues/736) is a discussion regarding to this issue that seems still there. 

After the installation, these steps help you get started right away. (Note that it assumes your training data are images)

1. Put your training images in $Caffe_root/data/yourdata folder. 

	There can be 'train', 'test', or 'val' folders for training, testing, and validation purposes. 

2. Prepare your image labels in a txt file in $Caffe\_root/data/yourdata folder. 

	Its a simple format like the following

	image\_01.jpg 0 

	image\_02.jpg 0 

	image\_03.jpg 1

	image\_04.jpg 2

	image\_05.jpg 0

	image\_06.jpg 3

	...

	There can be 'train.txt', 'test.txt', 'val.txt' for different cases accordingly. 

3. Get your training data in LMDB format, which is the right format for caffe

	Caffe has demonstrated how to do this in its examples. For example, under the folder $Caffe\_root/example/imagenet/
	there are several files very helpful. create\_imagenet.sh helps you convert your data (several folders with training images)
	into the desired format. make\_imagenet\_mean.sh helps you to extract the mean of images, which helps improving the training
	performance. 

4. Customize your solver.prototxt and train\_test.prototxt.  

	solver.prototxt defines the optimizers, loss functions, number of iterations, etc. 

	train\_test.prototxt defines the topology of the convolutional neural network. 

5. Start training! 

	Take a look at $Caffe\_root/example/imagenet/train\_caffenet.sh and you will know how to start training.

6. Scoring

	After the training is done, you can test your model on some images and see how it predicts. 

	To do that, check out $Caffe\_root/build/examples/cpp\_classification/classification.bin
