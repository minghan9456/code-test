<?php
$a = [1, 5, 3, 3, 7];
$a = [1, 3, 5];

/*
$a = [1, 0, 0, 1, 0, 0, 0,0,0,0,0];
$a = [0, 0, 1, 1, 0, 0];
*/

$r = solution($a);
var_dump($r);


// test 3
function solution($A) {

	$holder = $A;
	sort($holder);

	$assocRet = array_intersect_assoc($holder, $A);
	$check = count($A) - count($assocRet);

	return ($check == 2 || $check == 0) ? true : false;
}

/*
//test 2
function solution($A) {
    // write your code in PHP7.0
    return min(array_count_values($A));
}
*/

/*
//test 1
function solution($A) {
    // write your code in PHP7.0
    $ret = 0;
    
    foreach($A as $val) {
        preg_match('/^(?:-?\d{2})?$/', $val, $matches, PREG_OFFSET_CAPTURE);
        if ($matches) {
            $ret += $val;
        }
    }
    
    return $ret;
}
*/

/*
//test 4
-- write your code in PostgreSQL 9.4
SELECT dept_id, count(emp_id) as count, sum(salary) as sum_of_salary FROM employee 
GROUP BY employee.dept_id 
ORDER BY employee.dept_id
*/
