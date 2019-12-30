<?php

$matrix = [
  [0, 1, 1, 2], 
  [0, 5, 0, 0], 
  [2, 0, 3, 3]
];
$r = matrixElementsSum($matrix);
var_dump($r);

$matrix = [
  [1, 1, 1, 0], 
  [0, 5, 0, 1], 
  [2, 1, 3, 10]
];
$r = matrixElementsSum($matrix);
var_dump($r);

function matrixElementsSum($matrix) {
  $sum = 0;

  for($i = 0; $i <= count($matrix[0]) - 1; $i++) {
    for($j = 0; $j <= count($matrix) - 1; $j++) {
      if ($matrix[$j][$i] == 0) {
        break;
      }
      $sum += $matrix[$j][$i];
    }
  }

  return $sum;
}
