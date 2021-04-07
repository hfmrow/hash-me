## Informations about

## Hash Me

At the bottom you can find a compiled standalone ".deb" version with its checksum. The ".tar.gz" & ".zip" sources contain a "vendor" directory ensuring you can always compile it even if the official libraries have been changed.

## Changelog

All notable changes to this project will be documented in this file.

## [1.2] 2021-04-08

### Added

- Concurrent files processing (**gui** and **cli**).
- Tasks are processed in independent processes and no longer block the primary interface during processing, even if non-concurrent processing is chosen (**gui**).
- Recursive files scanner capability (**gui**), already existing with (**cli**) version.

### Changes

- The hash calculation is no longer calculated twice (Continue button), computation is only performed when the hash files / methods change.

- Some code rewriting for more reliability (**gui** and **cli**).

---

## [1.1] 2021-04-02

- First public release