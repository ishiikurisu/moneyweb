/**
 * Calls a function when the page is ready.
 */
function ready(fn) {
  if (document.readyState != 'loading'){
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

/**
 * Loads the log from memory
 */
function loadUserDatabase(callback) {
    return callback({
        error: "not implemented yet",
        log: null
    });
}
