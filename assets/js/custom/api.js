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
            let url = base_url + "/api/service-transactions-complex";
            console.log(customer);
            console.log(service);
            console.log(additionalItem);

            var payload = {
                customer: customer,
                serviceTransaction: service,
                additionalItems: additionalItem,
            };
            console.log(payload);

            $.ajax({
                url: url,
                method: 'POST',
                data: JSON.stringify(payload),
                contentType: 'application/json',
                success: callback
            });
        },
        ServiceTransactionUpdate: function(payload, id, callback) {
            let url = base_url + "/api/service-transactions-complex/" + id;

            $.ajax({
                url: url,
                method: 'PUT',
                data: JSON.stringify(payload),
                contentType: 'application/json',
                success: callback
            });
        },
        ServiceTransactionDelete: function(id, callback) {
            let url = base_url + "/api/service-transactions/" + id;

            $.ajax({
                url: url,
                method: 'DELETE',
                contentType: 'application/json',
                success: callback,
            });
        },
        ServiceTransactionSendStatusEmail: function(id, callback) {
            let url = base_url + "/api/service-transactions/sendstatusemail/" + id;

            $.ajax({
                url: url,
                method: 'GET',
                contentType: 'application/json',
                success: callback,
            });
        },
        init: function() {

        }
    };
});