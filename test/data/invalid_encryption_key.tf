provider "instaclustr" {
    username = "%s"
    api_key = "%s"
}

resource "instaclustr_encryption_key" "invalid" {
<<<<<<< HEAD
    alias = "%s"
=======
    alias = "ic_test_key"
>>>>>>> e2a4bb19800c323c205c06f49c37775b3319210e
    arn = "%s!@#$"
}
