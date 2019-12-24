<?php
$nums = array(1, 2, 3, 4, 5, 6, 7, 8, 9);

foreach($nums as $key => $num) {
	foreach($nums as $_num) {
		$res = $_num * $num; 
		print_r("{$_num}*{$num}={$res}\t");
	}
	print_r("\n");
}
