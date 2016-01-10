#include <libusb.h>
#include <stdio.h>
#include <string.h>

#include "names.h"

const int ERR_FAILED_GET_DEVICE_LIST = -1;
const int ERR_FIND_NO_MATCHED_DEV = -2;

// FindFirstMatchedProduct finds the first found USB device whose
// product name contains subs and returns 0.  It returns a negative
// number for errors.
int FindFirstMatchedProduct(libusb_context *ctx,
			    const char* subs,
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

    char vendor[128];
    get_vendor_string(vendor, sizeof(vendor),
		      desc.idVendor);
    
    if (strstr(product, subs) != NULL or strstr(vendor, subs) != NULL) {
      *found = list[i];
      return 0;
    }
  }
  return ERR_FIND_NO_MATCHED_DEV;
}

int main() {
  if (names_init() < 0) {
    return -1;
  }
  
  libusb_context *ctx = NULL;
  int r = libusb_init(&ctx);
  if (r < 0) {
    fprintf(stderr, "USB initialization error: %d\n", r);
    return -1;
  }

  libusb_device* arduino = NULL;
  r = FindFirstMatchedProduct(ctx, "Arduino", &arduino);
  if (r != 0) {
      printf("Not found: %d\n", r);
  }

  if (arduino != NULL) {
    printf("Found it!\n");
  }
  
  libusb_exit(ctx);
  names_exit();
  
  return 0;
}

