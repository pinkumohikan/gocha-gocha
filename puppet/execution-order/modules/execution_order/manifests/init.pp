class execution_order::install {
  notify {
    "notify_install":
      message => "install"
  }
}

class execution_order::prepare {
  notify {
    "notify_prepare":
      message => "prepare"
  }
}

class execution_order {
  include execution_order::install
  include execution_order::prepare

  notify {
    "notify_hoge_one":
      message => "hoge_one"
  }

  Class["execution_order::install"]
  -> Class["execution_order::prepare"]
  -> Class["execution_order"]

  notify {
    "notify_hoge_two":
      message => "hoge_two"
  }
}
