# Changes

## 1.0.0

Major changes on frontend and backend. The API contains breaking changes, but neither the Go nor the PHP client are affected by these changes.

### Major changes

* Switched frontend to React.
* Switched backend to an event-sourcing architecture (migration is automatic).
* API no longer returns numeric IDs, but uses the name of environments and features for identity.

### New features

* Features can be deleted from the UI.
* Always migrate database on startup.
* Increased creation timestamp precision from days to seconds.
* Order environments by creation time in the UI.
* Add API endpoint to order environments manually.
* Auto focus input fields in the UI.

### Bugfixes and minor changes

* Unknown API routes now properly return 404.
* Unknown frontend routes now redirect to the home page.

## 0.8.0

This is the first versioned release.

### New features

* Features are now sorted in reverse chronological order.

### Bugfixes and minor changes

* Names with spaces, slashes, and various other characters no longer break the UI.
* Fix typescript error with moment.js.

## 0.0.0

This corresponds to everything up to commit `3637a6e` (inclusive) made on Jan 29, 2018.
