open Bsearch_benchmark
open OUnit2

let test_tree =
  let rec helper t = function
    | 5 -> helper t (5 + 1)
    | 10 -> t
    | n -> helper (insert n t) (n + 1)
  in
  helper Leaf 0

let make_test name expected input =
  name >:: fun _ -> assert_equal expected input

let tests =
  "test suite for binary tree"
  >::: [
         make_test "onecomp present" true (bsearch_onecomp 2 test_tree);
         make_test "twocomp present" true (bsearch_twocomp 2 test_tree);
         make_test "onecomp too low" false (bsearch_onecomp (-2) test_tree);
         make_test "twocomp too low" false (bsearch_twocomp (-2) test_tree);
         make_test "onecomp too high" false (bsearch_onecomp 12 test_tree);
         make_test "twocomp too high" false (bsearch_twocomp 12 test_tree);
         make_test "onecomp not mid" false (bsearch_onecomp 5 test_tree);
         make_test "twocomp not mid" false (bsearch_twocomp 5 test_tree);
       ]

let _ = run_test_tt_main tests
