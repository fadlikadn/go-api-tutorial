"use strict";

function uniqueID() {
    return '_' + Math.random().toString(36).substr(2, 9);
}

$(document).ready(function() {
    // To fix issue modal disable scroll when open modal on top of modal
    $('.modal').on('hidden.bs.modal', function(e) {
        if ($('.modal:visible').length) {
            $('body').addClass('modal-open');
        }
    });
});