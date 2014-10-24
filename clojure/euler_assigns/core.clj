(ns euler-assigns.core
  (:gen-class))

(defn problem2 []
  (defn fib [prev cur sum]
    (cond (< cur 4000000)
          (cond (= 0 (mod cur 2))
                (fib cur (+ cur prev) (+ sum cur))
                true (fib cur (+ cur prev) sum))
           true sum))
  (fib 1 2 0))

; Suboptimal, but fun.
(defn problem4 []
  (defn palindrome? [n]
    (let [s (str n)]
      (= s (clojure.string/reverse s))))
  (apply max
    (filter some?
      (for [i (range 100 1000) j (range i 1000)]
        (let [prod (* i j)]
          (cond (palindrome? prod) prod))))))


(defn test-problem [prob desired]
  (let [fname (clojure.string/replace
              (str prob) #"^.+?\$([\w-]+?)@.+$" "$1")]
    (cond (= desired (prob)) (printf "%s passed!\n" fname)
          true (printf "%s → %d, expected %d\n"
                        fname (prob) desired))))

(defn -main
  "Some Project Euler problems."
  [& args]
  (test-problem problem2 4613732)
  (test-problem problem4 906609))
