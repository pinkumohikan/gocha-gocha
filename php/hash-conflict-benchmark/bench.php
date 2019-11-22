<?php

function main(array $argv): void
{
    $length = $argv[1];
    $hashs = generateHashs($length);

    var_dump('count($hashs):', count($hashs));
    var_dump('count(array_unique($hashs)):', count(array_unique($hashs)));

    var_dump('memory_get_peak_usage(MB):', memory_get_peak_usage(true) / 1024 / 1024);
}

function generateHashs(int $max): array
{
    $hashs = [];

    for ($i = 1; $i <= $max; $i++) {
        $hashs []= hash('sha256', $i);
    }

    return $hashs;
}

main($argv);
