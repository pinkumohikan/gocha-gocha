# ファイルへのパスを求める
get_path <- function() {
    paths <- c()

    for (fn in 0:(sys.nframe() - 1)) {
        paths <- append(paths, sys.frame(fn)$ofile)
    }

    return(tail(na.omit(paths), n=1))
}
