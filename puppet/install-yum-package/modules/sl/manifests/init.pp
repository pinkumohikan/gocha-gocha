class sl::install {
    package { 'sl':
        provider        => yum,
        ensure          => latest,
        install_options => ['--enablerepo=epel'],
        allow_virtual   => true
    }

    # 戦争起きないかな (ニコニコ)
    package { 'emacs':
        provider        => yum,
        ensure          => latest,
        allow_virtual   => true,
        require         => Package['vim']
    }
}

class sl::prepare {
}

class sl {
    include sl::install
    include sl::prepare

    Class["sl::install"]
    -> Class["sl::prepare"]
    -> Class["sl"]
}
