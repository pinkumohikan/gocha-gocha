# ファイルへのパスを求める
get_path <- function() {
    paths <- c()

    for (fn in 1:sys.nframe()) {
        path <- sys.frame(fn)$ofile
        if (!is.null(path)) {
            paths <- append(paths, path)
        }
    }

    return(tail(paths, n=1))
}
