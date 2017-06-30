options(java.parameters = "-Xmx5g")

library(RJDBC, quietly = TRUE)

connectRedshift <- function() {
    # see
    #   * http://docs.aws.amazon.com/ja_jp/redshift/latest/mgmt/connecting-ssl-support.html#connect-using-ssl
    #   * http://docs.aws.amazon.com/ja_jp/redshift/latest/mgmt/configure-jdbc-options.html

    driver <- JDBC('com.amazon.redshift.jdbc.Driver', '/RedshiftJDBC42-1.2.1.1001.jar', identifier.quote="`")
    url <- sprintf(
        'jdbc:redshift://%s:%s/%s?AuthMech=REQUIRE&ssl=true&sslMode=verify-ca&sslrootcert=%s&user=%s&password=%s',
        Sys.getenv('REDSHIFT_HOST'),
        Sys.getenv('REDSHIFT_PORT'),
        Sys.getenv('REDSHIFT_DB'),
        './redshift/redshift-ssl-ca-cert.pem',
        Sys.getenv('REDSHIFT_USER'),
        Sys.getenv('REDSHIFT_PASS')
    )
    return(dbConnect(driver, url))
}
