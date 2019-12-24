<?php

$matrix = [
  [1, 0, 0], 
  [0, 5, 0], 
  [0, 0, 3]
];
// true

$matrix = [
  [0, 1, 1], 
  [1, 0, 1], 
  [1, 1, 0]
];
// false

$r = isDiagonalMatrix($matrix);
var_dump($r);

function isDiagonalMatrix($matrix) {
  $size = count($matrix[0]);

  for ($row = 0; $row <= $size; $row++) {
    for ($column = 0; $column <= $size; $column++) {
      if ($row != $column) {
        if ($matrix[$row][$column] || $matrix[$column][$row] || $matrix[$row][$column] != $matrix[$column][$row]) return false;
      }
    }
  }

  return true;
}

