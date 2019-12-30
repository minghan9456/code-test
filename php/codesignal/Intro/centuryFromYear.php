<?php

$r = centuryFromYear(2009);
var_dump($r);

$r = centuryFromYear(9);
var_dump($r);

function centuryFromYear($year) {
  $century = 0;
  while($year > 0) {
    $century++;
    $year = $year - 100;
  }

  return $century;
}

