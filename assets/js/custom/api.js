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
        CustomerAddEdit: function(url, method, payload, callback) {
            $.ajax({
                url: url,
                method: method,
                data: JSON.stringify(payload),
                contentType: 'application/json',
                success: callback
            });
        },
        StoreNewServiceTransaction: function(customer, service, additionalItem, callback) {

        },
        init: function() {

        }
    };
});