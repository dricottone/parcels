# parcels

A tool and library to *parse URLs*.
If you say that fast enough, it sounds like *parcels*.

This is mainly intended for use in mail filters.


## Usage

`cat EMAIL | parcels`

This will...

 1. re-prints `EMAIL` with URLs replaced by indices
 2. print a postscript of indexed URLs

`cat EMAIL | parcels -n 0`

This will...

 1. try to find the first URL in `EMAIL` and re-print just that URL
 2. otherwise print nothing


## License

The URL regular expression is adapted from that used in
[urlscan](https://github.com/firecat53/urlscan), which is licensed under GPL v2.
Credit for this portion of the code should go to:

 + Scott Hansen \<[firecat4153@gmail.com](mailto:firecat4153@gmail.com)\>
   (Author and Maintainer)
 + Maxime Chatelle \<[xakz@rxsoft.eu](mailto:xakz@rxsoft.eu)\>
   (Debian Maintainer)
 + Daniel Burrows \<[dburrows@debian.org](mailto:dburrows@debian.org)\>
   (Original Author)

I personally consider this project to be derivative, therefore I am keeping the
GPL v2 license of the urlscan project.

