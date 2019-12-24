<?php

$r = makeArrayConsecutive2([6, 2, 3, 8]);
var_dump($r);

function makeArrayConsecutive2($statues) {
  $max = $min = $statues[0];

  for ($i = 1; $i <= count($statues) -1; $i++) {
    if ($statues[$i] > $max) $max = $statues[$i];
    if ($statues[$i] < $min) $min = $statues[$i];
  }

  $count = 0;
  for ($i = $min; $i <= $max; $i++) {
    if (!in_array($i, $statues)) $count++;
  }

  return $count;
}
