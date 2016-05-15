
int H = 64, W = 64;
void setup() {
  size(64, 64);  // image dimension
  stroke(250);     // line color
  background(243, 206, 3);  // the yellowiesh color
  noLoop();
}

int getSign() {
  return random(1) < 0.5 ? -1 : 1;
}

String getFileLoc(int folder, int index) {
  String loc = str(folder) + '/' + str(index) + ".jpg";
  return loc;
}

void draw() {
  float d = 32, l_width = 5, theta;
  int N = 50000;  // total number of image
  int[] theta_index = new int[11];
  for (int i = 0; i < 11; i++) {
    theta_index[i] = 15 * (i+1);
  }
  for (int i = 0; i < N; ++i) {
    int ti = int(randomGaussian() * 2 + 5);
    if (ti < 0 || ti > 10) {
      continue;
    }
    theta = theta_index[ti] * PI / 180;
    int sign = getSign();
    background(243 + random(20) * sign, 206 + random(20) * sign, 3);  
    String f_loc = getFileLoc(theta_index[ti], i);
    imageGen(d + random(10) * sign, theta + random(5) * sign * PI / 180, l_width, f_loc);
  }
}

float[] linepoints(float x1, float theta) {
  float[] res = new float[2];
  float theta1 = atan(H / (W-x1)), theta2 = PI-atan(H / x1);
  float x, y;
  if (theta < theta1) {
    y = (W - x1) * tan(theta);
    x = W;
  } else if (theta > theta1 && theta < theta2) {
    y = H;
    x = x1 + H / tan(theta);
  } else {
    x = 0;
    y = -tan(theta) * x1;
  }
  res[0] = x;
  res[1] = y;
  return res;
}


void imageGen(float d, float theta, float l_width, String file_loc) {
  strokeWeight(l_width);
  float x1 = (W-d)/2, x2 = (W+d)/2;
  float xoff = noise(1) * x1;
  int sign = getSign();
  x1 -= xoff * sign;
  x2 -= xoff * sign;
  float[] xy = linepoints(x1, theta);
  line(x1, 0, xy[0], xy[1]);
  xy = linepoints(x2, theta);
  line(x2, 0, xy[0], xy[1]);
  save(file_loc);
}
