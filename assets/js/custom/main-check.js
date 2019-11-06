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
                        $('.damage-type').html(data.damage_type);
                        $('.equipment').html(data.equipment);
                        $('.description').html(data.description);
                        $('.technician').html(data.technician);
                        $('.repair-type').html(data.repair_type);
                        $('.spare-part').html(data.spare_part);
                        $('.total-price').html(accounting.formatMoney(data.total_price, "Rp ", 0, ".", ","));
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