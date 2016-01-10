#include <libusb.h>
#include <stdio.h>
#include <string.h>

#include "names.h"

const int ERR_FAILED_GET_DEVICE_LIST = -1;
const int ERR_FIND_NO_MATCHED_DEV = -2;

// FindFirstMatchedProduct finds the first found USB device whose
// product name contains product_subs and returns 0.  It returns a
// negative number for errors.
int FindFirstMatchedProduct(libusb_context *ctx,
			    const char* product_subs,
			    libusb_device** found) {
  libusb_device **list;	  
  int num_devs = libusb_get_device_list(ctx, &list);
  if (num_devs < 0) {
    return ERR_FAILED_GET_DEVICE_LIST;
  }

  for (int i = 0; i < num_devs; i++) {
    struct libusb_device_descriptor desc;
    libusb_get_device_descriptor(list[i], &desc);

    char product[128];
    get_product_string(product, sizeof(product),
		       desc.idVendor,
		       desc.idProduct);

    if (strstr(product, product_subs) != NULL) {
      *found = list[i];
      return 0;
    }
  }
  return ERR_FIND_NO_MATCHED_DEV;
}

int main() {
  libusb_context *ctx = NULL;
  int r = libusb_init(&ctx);
  if (r < 0) {
    fprintf(stderr, "USB initialization error: %d\n", r);
    return -1;
  }

  libusb_device* arduino;
  FindFirstMatchedProduct(ctx, "Arduino", &arduino);

  libusb_exit(ctx);
  return 0;
}

