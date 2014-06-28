;;; theme-creator.el --- theme creator in emacs lisp.

;;; Code:

(eval-when-compile (require 'cl-lib))
(require 'widget)

(defvar dark-background nil)

(defgroup theme-creator nil
  "Create emacs24 themes from within emacs"
  :group 'theme-creator)

(defcustom theme-creator-buffer-name "*theme-creator*"
  "Name used for theme-creator buffer."
  :group 'theme-creator
  :type 'string)

(defvar  *theme-name*  "")

(defun set-theme-name (name)
  "set *theme-name* to provided name."
  (setf *theme-name* name))

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

(defun set-background (str)
  "set background to light or dark according to supplied string."
  (if (string=  str "Light")
      (set-light-background)
    (set-dark-background)))

(defun theme-creator ()
  ""
  (interactive)
  (select-window (or (get-buffer-window theme-creator-buffer-name)
                     (selected-window)))
  (switch-to-buffer theme-creator-buffer-name)
  (init-creator-grid))


(defun init-creator-grid ()
  (widget-insert "some documentation\n\n")
  (widget-create 'editable-field
                 :size 20
                 :notify (lambda (widget &rest ignore)
                           (set-theme-name (widget-value widget)))
                 :format "%v"; Text after field!
                 "Theme Name")
  (widget-insert "\n\n")
  (widget-create 'radio-button-choice
                  :value "Light"
                  :notify (lambda (widget &rest ignore)
                            (set-background (widget-value widget))
                            (message "You set the background to %s."
                                     (widget-value widget)))
                  '(item "Light") '(item "Dark"))
  (widget-insert "\n")
  (widget-setup)
  (use-local-map widget-keymap))















(provide 'theme-creator)

;;; theme-creator.el ends here
