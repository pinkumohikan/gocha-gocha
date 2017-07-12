<?php
require_once __DIR__.'/../bootstrap.php';

$app = new Slim\App();

$app->get('/', function ($request, $response, $args) {
    return $response->withJson(['now' => (new \DateTimeImmutable())->format('Y-m-d H:i:s')]);
});

$app->run();

