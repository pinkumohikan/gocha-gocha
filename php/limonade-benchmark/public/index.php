<?php

require dirname(__DIR__).'/bootstrap.php';

dispatch('/', function () {
    echo json_encode(['now' => (new \DateTimeImmutable())->format('Y-m-d H:i:s')]);
});

run();
