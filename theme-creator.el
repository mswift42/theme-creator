;;; theme-creator.el --- theme creator in emacs lisp.

;;; Code:

(eval-when-compile (require 'cl-lib))

(defvar dark-background nil)

(defgroup theme-creator nil
  "Create emacs24 themes from within emacs"
  :group 'theme-creator)

(defcustom theme-creator-buffer-name "*theme-creator*"
  "Name used for theme-creator buffer."
  :group 'theme-creator
  :type 'string)

(make-variable-buffer-local 'dark-background)

(defun dark-background-p ()
  "is background set to dark?"
  dark-background)

(defun set-dark-background ()
  "set background to dark"
  (setf dark-background t))

(defun set-light-background ()
  "set background to light"
  (setf dark-background nil))

(defun theme-creator ()
  ""
  (interactive)
  (select-window (or (get-buffer-window theme-creator-buffer-name)
                     (selected-window)))
  (switch-to-buffer theme-creator-buffer-name))















(provide 'theme-creator)

;;; theme-creator.el ends here
