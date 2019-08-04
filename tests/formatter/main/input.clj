(if 1 2 3)

(if 1 2 )

(if 1)

(if)

(fn [a b] 1 2)

(fn a [a b] 1 2)

(let [a 1 b 2] 1 2)

(let [] 1)

(let [a 1] 1 2)

(let [a] 1 2)

(loop [a 1 b 1] (if (> a 10) [a b] (recur (inc a) (dec b))))

(letfn [(neven? [n] (if (zero? n) true (nodd? (dec n))))
        (nodd? [n] (if (zero? n) false (neven? (dec n))))]
  (neven? 10))

(do 1 2)

(try 1 (catch Exception e 2) (finally 3))

(defn plus [x y] (+ x y))

(defn cast
  "Throws an error if x is not of a type t, else returns x."
  {:added "1.0"}
  [t x]
  (cast__ t x))

(deftest t (is (= 1 2)))

(def PI 3.14)

(ns my.test
  (:require my.test1
            [my.test2]
            [my.test3 :as test3 :refer [f1]])
  (:import (java.time LocalDateTime ZonedDateTime ZoneId)
           java.time.format.DateTimeFormatter))

(defn test-docstring
  "Given a multimethod and a dispatch value, returns the dispatch fn
  that would apply to that value, or nil if none apply and no default"
  {:added "1.0"}
  [t]
  (print "ha
         ha"))

(test-call 1 2 3)

(test-call 1
           2
           3)

(test-call 1 2
           3)

(test-call
 1
 2
 3)

(test-call
 1 2
 3)

@mfatom

@1

#"t"

#'t

#^:t []

#(inc %)

'test

'[test]

`(if-not ~test ~then nil)

(defmacro and
  "Evaluates exprs one at a time, from left to right. If a form
  returns logical false (nil or false), and returns that value and
  doesn't evaluate any of the other expressions, otherwise it returns
  the value of the last expr. (and) returns true."
  {:added "1.0"}
  ([] true)
  ([x] x)
  ([x & next]
   `(let [and# ~x]
      (if and# (and ~@next) and#))))

(def
  ^{:arglists '([& items])
    :doc "Creates a new list containing the items."
    :added "1.0"
    :tag List}
  list list__)

{1 2 3 4}

{1 2
 3 4}

[1 2
 3 4]

#{1 2
  3 4 5}

[#?(:cljs 1)]
(#?(:cljs 1))
#?(:clj 1)
#?@(:cljs 3)
(def regexp #?(:clj re-pattern :cljs js/XRegExp))


;; Should FAIL

#?(:cljs)
#?(:cljs (let [] 1) :default (let [] 1))
[#?@(:clj 1)]

#:t{:g 1}
#::{:g 1}
#:t{:_/g 1}
#:t{:h/g 1}
#::s{:g 1}

;; Should FAIL
#::{g 1}

#inst 1
#uuid 2

;; Should FAIL

#t 4
#g [a]

(defn ^:private line-seq*
  [rdr]
  (when-let [line (reader-read-line__ rdr)]
    (cons line (lazy-seq (line-seq* rdr)))))

