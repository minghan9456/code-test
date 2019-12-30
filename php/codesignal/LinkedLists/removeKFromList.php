<?php

/*
$l = new ListNode(3);
$l->next = new ListNode(1);
$l->next->next = new ListNode(2);
$l->next->next->next = new ListNode(3);
$l->next->next->next->next = new ListNode(4);
$l->next->next->next->next->next = new ListNode(5);
//print_r($l);
$k = 3;

$r = removeKFromList($l, $k);
print_r($r);
*/

$r = removeKFromList($l, -1000);
print_r($r);


// Singly-linked lists are already defined with this interface:
class ListNode {
  public $value;
  public $next;

  public function __construct($x) {
    $this->value = $x;
    $this->next = null;
  }
}

function removeKFromList($l, $k) {
  if ($l === null) {
    return $l;
  }

  if ($l->value === $k) {
    if ($l->next) {
      return removeKFromList($l->next, $k);
    }
  }
  else {
    $tmp = new ListNode($l->value);
    if ($l->next) {
      $tmp->next = removeKFromList($l->next, $k);
    }
    return $tmp;
  }

  return;
}

