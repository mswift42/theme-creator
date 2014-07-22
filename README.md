emacs-theme-creator
=============

create emacs 24 themes.

Installation Instructions:
--------------------------

No installation necessary. Emacs Theme Creator runs on Appengine, using Go and a bit of AngularJS.

Usage:
------

Create a new directory for your theme to live in.

Create an empty theme file in the new directory:

    touch <name of theme>-theme.el

It is highly recommended that you store this in some form of VCS.

If your new theme is called `rainforest`, your theme file has be named `rainforest-theme.el`. This is important if you are planning to submit your theme to melpa.

Go to [emacs-theme-creator](http://emacs-theme-creator.appspot.com/).

Enter theme name and author name, and if you have your new theme in a VCS repo, it's url.

You need to submit a color for every face listed, either by entering a color in hex format, i.e. `#202020`, or by clicking on the form using the supplied colorpicker.

When you are finished, click "Done".

Select all: `C-a`, copy: `C-c`, and save in your earlier created file.


Applying a new theme
--------------------

Disable the theme you are using at present.

    (disable-theme <theme-in-use>)

Try out the new theme.

    (load-file <your-name-theme-file.el>)

Customize the theme
-------------------

Visit your new theme.

    `C-x-C-f` <filename.el>

Recommended:

load rainbow-mode to display colors for the supplied colors.

Customize to your likings.

You can see a list of all faces and their current value with

    `M-x list-faces-display`


Used Libraries
---------------

Emacs-theme-creator uses Go, [AngularJS](https://angularjs.org/), in particular [angular-seed])(https://github.com/angular/angular-seed) for the live preview of the selected faces and  [go-colorful](https://github.com/lucasb-eyer/go-colorful) for calculating darker and lighter variants of selected faces.

Screenshots
-----------

picking faces:

![Screenshot](https://github.com/mswift42/theme-creator/raw/master/screen1.png)

theme file:

![Screenshot](https://github.com/mswift42/theme-creator/raw/master/screen2.png)


TODO:
-----

Add more preview languages.


Donate:

Every little helps:


[![PayPayl donate button](http://img.shields.io/paypal/donate.png?color=yellow)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=5823VL6B3XM86 "Donate once-off to this project using Paypal")
