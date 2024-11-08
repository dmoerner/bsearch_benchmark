type 'a tree = Leaf | Node of 'a * 'a tree * 'a tree

let rec insert n = function
  | Leaf -> Node (n, Leaf, Leaf)
  | Node (v, l, r) ->
      if n < v then Node (v, insert n l, r)
      else if n > v then Node (v, l, insert n r)
      else Node (v, l, r)

let rec bsearch_twocomp n = function
  | Leaf -> false
  | Node (v, l, r) ->
      if n < v then bsearch_twocomp n l
      else if n > v then bsearch_twocomp n r
      else true

let bsearch_onecomp n t =
  let rec bsearch_helper target t cand =
    match (t, cand) with
    | Leaf, None -> false
    | Leaf, Some v -> v = target
    | Node (v, l, r), c ->
        if target < v then bsearch_helper target l c
        else bsearch_helper target r (Some v)
  in
  bsearch_helper n t None

(* Worst case is an unbalanced tree where n > v *)

let build_tree size =
  let rec helper t = function
    | 0 -> t
    | n -> helper (insert (Random.int size) t) (n - 1)
  in
  helper Leaf size
