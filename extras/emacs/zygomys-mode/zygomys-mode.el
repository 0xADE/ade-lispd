;; Define zygomys-mode
(define-derived-mode zygomys-mode lisp-mode "zygomys"
  "Major mode for editing Zygomys Lisp code with Go-style comments."
  :syntax-table zygomys-mode-syntax-table

  ;; Set up font-lock keywords for syntax highlighting
  (setq font-lock-defaults '((zygomys-font-lock-keywords) nil t))

  ;; Add Zygomys-specific indentation rules if needed
  (setq indent-line-function 'lisp-indent-line)

  ;; Enable electric-pair-mode for automatic matching of parentheses
  (electric-pair-mode 1)

  ;; Add Zygomys-specific keybindings if needed
  (define-key zygomys-mode-map (kbd "C-c C-c") 'zygomys-eval-region)
  (define-key zygomys-mode-map (kbd "C-c C-l") 'zygomys-load-file))

;; Define a custom syntax table for Zygomys mode
(defvar zygomys-mode-syntax-table
  (let ((table (make-syntax-table lisp-mode-syntax-table)))
    ;; Treat `//` as starting a single-line comment
    (modify-syntax-entry ?/ ". 124b" table)
    (modify-syntax-entry ?\n "> b" table)
    ;; Treat `/*` and `*/` as starting and ending multi-line comments
    (modify-syntax-entry ?* ". 23" table)
    table)
  "Syntax table for `zygomys-mode'.")

;; Define font-lock keywords for Zygomys Lisp
(defconst zygomys-font-lock-keywords
  `(
    ;; Highlight keywords
    ("\\<\\(def\\|let\\|if\\|fn\\|defn\\|do\\|for\\|quote\\|quasiquote\\|unquote\\|unquote-splicing\\)\\>" . font-lock-keyword-face)

    ;; Highlight built-in functions
    ("\\<\\(assert\\|map\\|filter\\|reduce\\|apply\\|cons\\|car\\|cdr\\|list\\)\\>" . font-lock-builtin-face)

    ;; Highlight constants
    ("\\<\\(true\\|false\\|nil\\)\\>" . font-lock-constant-face)

    ;; Highlight strings
    ("\"[^\"]*\"" . font-lock-string-face)

    ;; Highlight single-line comments
    ("//.*" . font-lock-comment-face)

    ;; Highlight multi-line comments
    ("/\\*[^*]*\\*+([^*/][^*]*\\*+)*/" . font-lock-comment-face)

    ;; Highlight numbers
    ("\\b[0-9]+\\b" . font-lock-number-face)
    ))

;; Function to evaluate a region of Zygomys code
(defun zygomys-eval-region (start end)
  "Evaluate the Zygomys code in the region from START to END."
  (interactive "r")
  (message "Evaluating Zygomys code... Not implemented yet."))

;; Function to load a Zygomys file
(defun zygomys-load-file ()
  "Load the current Zygomys file."
  (interactive)
  (message "Loading Zygomys file... Not implemented yet."))

;; Add Zygomys mode to auto-mode-alist
(add-to-list 'auto-mode-alist '("\\.zy\\'" . zygomys-mode))

(provide 'zygomys-mode)
