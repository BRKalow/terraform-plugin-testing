## 1.1.0 (February 06, 2023)

FEATURES:

* helper/resource: Added `TF_ACC_PERSIST_WORKING_DIR` environment variable to allow persisting of Terraform files generated during each test step ([#18](https://github.com/hashicorp/terraform-plugin-testing/issues/18))
* helper/resource: Added `TestCase` type `WorkingDir` field to allow specifying the base directory where testing files used by the testing module are generated ([#18](https://github.com/hashicorp/terraform-plugin-testing/issues/18))

## 1.0.0 (January 10, 2023)

NOTES:

* Same testing functionality as that of terraform-plugin-sdk v2.24.1, repacked in a standalone repository ([#24](https://github.com/hashicorp/terraform-plugin-testing/issues/24))

