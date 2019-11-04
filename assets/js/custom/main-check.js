"use strict";

var MainCheck = {};
var base_url = 'http://localhost:8080';

$(function() {
    MainCheck = {
        _handleButtons: function() {
            $(document).on('click', '#btnSearchInvoice', function(e) {
                e.preventDefault();
                let url = base_url + '/api/search/invoice/' + $('#invoice-no').val().toString();
                $.ajax({
                    url: url,
                    method: 'GET',
                    contentType: 'application/json',
                    success: function(data) {
                        console.log(data);
                        $('#invoiceFoundModal').modal('show');

                        $('.item-name').html(data.item_name);
                    },
                    error: function(err) {
                        console.log(err);
                        alert('invoice tidak ditemukan');
                    }
                });
            });
        },
        init: function() {
            this._handleButtons();
        }
    };
    MainCheck.init();
});