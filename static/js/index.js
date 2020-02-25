/**
 * returns the inner HTML for the logs element based on the log result object
 * @param result an object containing two entries:
 *               - error: null if the log was correctly downloaded, or a string
 *                        indicating the error
 *               - log: an array of entries if the log was correctly
 *                      downloaded, null otherwise
 * @returns the inner HTML content that best describes the results.
 */
function drawLogsContent(inlet) {
    var outlet = "";

    if (inlet.error) {
        outlet = `<p class="subtitle">Oops :( Try again later</p>`;
    } else if (inlet.log.length === 0) {
        outlet = `<p class="subtitle">Please add your first entry so we can get started!</p>`;
    } else {
        // TODO a bulma table using the log
        outlet = `NOT IMPLEMENTED YET!`;
    }

    return outlet;
}

/**
 * Called when the page first loads. Should either display the log or tell a
 * instructing the user to add his first entry.
 */
function main() {
    loadUserDatabase(function(result) {
        document.getElementById('logs').innerHTML = drawLogsContent(result);
    });
}
