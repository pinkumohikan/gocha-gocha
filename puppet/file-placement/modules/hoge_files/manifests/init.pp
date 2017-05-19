class hoge_files::install {
    file {
        "/tmp/hoge":
        mode    => 664,
        recurse => "remote",
        source  => "puppet:///modules/hoge_files/hoge/"
    }
}

class hoge_files::prepare {
}

class hoge_files {
    include "hoge_files::install"
    include "hoge_files::prepare"

    Class["hoge_files::install"]
    ->Class["hoge_files::prepare"]
    ->Class["hoge_files"]
}
