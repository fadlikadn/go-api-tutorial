"use strict";

const base_url = 'http://localhost:8080';
var APIs = {};

$(function() {
    APIs = {
        CustomerGetAll: function(callback) {
            var url = base_url + "/api/customers";
            $.ajax({
                url: url,
                method: 'GET',
                contentType: 'application/json',
                success: callback
            });
        },
        init: function() {

        }
    };
});