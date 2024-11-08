open Bsearch_benchmark
open Benchmark

let size = 5000
let tree = build_tree size
let random_bsearch_twocomp t = bsearch_twocomp (Random.int size) t
let random_bsearch_onecomp t = bsearch_onecomp (Random.int size) t

(*let resT =*)
(*  throughputN ~repeat:1 5*)
(*    [*)
(*      ("random_bsearch_twocomp", random_bsearch_twocomp, tree);*)
(*      ("random_bsearch_onecomp", random_bsearch_onecomp, tree);*)
(*    ]*)

let resL =
  latencyN ~repeat:1 10000000L
    [
      ("random_bsearch_twocomp", random_bsearch_twocomp, tree);
      ("random_bsearch_onecomp", random_bsearch_onecomp, tree);
    ]

(*let _ = resT*)
let _ = resL
