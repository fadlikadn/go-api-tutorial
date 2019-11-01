"use strict";

$(function() {
    ServiceTransactions = {
        customers: null,
        additionalItemsTableData: [],
        additionalCostTotal: 0,
        serviceCost: 0,
        allCostTotal: 0,
        selectedServiceTransactionData: {},
        _datatables: function() {
            var self = this;
            $('#table-service-transactions').DataTable({
                "processing": true,
                "searching": true,
                "serverSide": false,
                "paging": true,
                "bLengthChange": true,
                "ordering": true,
                "ajax": {
                    "url": base_url + "/api/service-transactions",
                    "dataSrc": "",
                    "data": function(d) {
                        console.log(d);
                    },
                    "error": function(d) {
                        console.log(d);
                    }
                },
                "columns": [
                    {
                        "data": "service_date", render: function(data, type, row, meta) {
                            return moment(data).format('DD/MM/YYYY');
                        }
                    },
                    {
                        "data": "invoice_no", "defaultContent": ""
                    },
                    {
                        "data": "customer", render: function(data, type, row, meta) {
                            return data.name;
                        }
                    },
                    {
                        "data": "item_name", "defaultContent": ""
                    },
                    {
                        "data": "damage_type", "defaultContent": ""
                    },
                    {
                        "data": "repair_type", "defaultContent": ""
                    },
                    // {
                    //     "data": "price", "defaultContent": ""
                    // },
                    {
                        "data": "total_price", render: function(data, type, row, meta) {
                            return `<span class="float-right">${accounting.formatMoney(data, "Rp ", 0, ".", ",")}</span>`;
                        }
                    },
                    {
                        "data": "status", "defaultContent": ""
                    },
                    // <a href='#' class='btn btn-sm btn-success tell_customer' data-key=${data} data-object='${JSON.stringify(row)}'>Tell Customer</a>
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#serviceTransactionEditModal' data-toggle="modal" class='btn btn-sm btn-success service-transaction_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#serviceTransactionDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger service-transaction_delete' data-key=${data} >Delete</a> &nbsp;
<div class="dropdown">
    <button class="btn btn-sm btn-success dropdown-toggle" type="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
        Tell Customer
    </button>
    <div class="dropdown-menu">
        <a class="dropdown-item" href="#" data-target="#whatsappConfirmationModal" data-toggle="modal" data-object='${JSON.stringify(row)}' aria-haspopup="true" aria-expanded="false">Whatsapp</a>
    </div>
</div>
`;
                        }
                    }
                ]
            });

            $('#table-additional-cost-modal')
                .on('preXhr.dt', function(e, settings, data) {
                    console.log('reset additionalCostTotal before table load new data');

                    ServiceTransactions.additionalCostTotal = 0;
                    ServiceTransactions.allCostTotal = 0;
                    ServiceTransactions.allCostTotal += ServiceTransactions.serviceCost;
                    self._calculateAllCostTotal();
                })
                .DataTable({
                "processing": true,
                "searching": true,
                "serverSide": false,
                "paging": true,
                "bLengthChange": true,
                "ordering": true,
                "ajax": function(data, callback, settings) {
                    callback({
                        data: (ServiceTransactions.additionalItemsTableData.length > 0) ? ServiceTransactions.additionalItemsTableData : [],
                    })
                },
                "columns": [
                    {
                        "data": "name", "defaultContent": ""
                    },
                    {
                        "data": "notes", "defaultContent": ""
                    },
                    {
                        "data": "cost", render: function(data, type, row, meta) {
                            return `<span class="float-right">${accounting.formatMoney(data, "Rp", 0, ".", ",")}</span>`;
                        }
                    },
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#additionalCostModal' data-toggle="modal" class='btn btn-sm btn-success additional-cost_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#additionalCostDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger additional-cost_delete' data-key=${data} >Delete</a>`;
                        }
                    }
                ],
                "rowCallback": function(row, data, index) {
                    ServiceTransactions.additionalCostTotal += parseInt(data.cost);
                    console.log(ServiceTransactions.additionalCostTotal);
                    self._calculateAllCostTotal();
                }
            });
        },
        _calculateAllCostTotal: function() {
            ServiceTransactions.allCostTotal = ServiceTransactions.serviceCost + ServiceTransactions.additionalCostTotal;
            $('#serviceTransactionEditModal #service-transaction-total-price').val(accounting.formatMoney(ServiceTransactions.allCostTotal, "", 0, ".", ","));
        },
        _APIs: function() {
            APIs.CustomerGetAll(function(res) {
                ServiceTransactions.customers = res;
                console.log(ServiceTransactions.customers);
            })
        },
        _datepicker: function() {
            console.log('set datepicker');
            $('#serviceTransactionEditModal #service-transaction-date').datetimepicker({
                format: 'DD/MM/YYYY',
            });

            $('#serviceTransactionEditModal #service-transaction-taken-date').datetimepicker({
                format: 'DD/MM/YYYY',
            });
        },
        _mapSelectedCustomer: function(customers, selectedCustomer) {
            var $serviceTransactionCustomer = $('#serviceTransactionEditModal #service-transaction-customer');
            $serviceTransactionCustomer.empty();
            $serviceTransactionCustomer.append($("<option/>"));
            $.each(customers, function(key, value) {
                let selected = false;
                if (value.id === selectedCustomer) selected = true;
                let $option = $("<option/>", {
                    value: value.id,
                    text: value.name,
                    selected: selected,
                    object: JSON.stringify(value)
                });
                $serviceTransactionCustomer.append($option);
            });
            $serviceTransactionCustomer.select2({
                placeholder: "Select Customer",
                allowClear: true
            });
        },
        _handleButtonEvents: function() {
            // Service Transaction Handle Button Events
            var self = this;
            var $serviceTransactionEditModal = $('#serviceTransactionEditModal');
            var $serviceTransactionDeleteModal = $('#serviceTransactionDeleteModal');
            var $whatsappConformationModal = $('#whatsappConfirmationModal');

            $whatsappConformationModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('object') != undefined) {
                    let transactionObject = $(e.relatedTarget).data('object');
                    let number = libphonenumber.parsePhoneNumberFromString(transactionObject.customer.phone, 'ID').number;
                    $('#whatsappConfirmationModal #customer-whatsapp-number').val(number);
                    $('#whatsappConfirmationModal #customer-object').val(JSON.stringify(transactionObject));
                }
            });

            $serviceTransactionDeleteModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Delete
                    var serviceTransactionId = $(e.relatedTarget).data('key');
                    $('#btnServiceTransactionDelete').attr('data-id', serviceTransactionId);
                }
            });

            $serviceTransactionEditModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Edit
                    var serviceTransactionId = $(e.relatedTarget).data('key');
                    var serviceTransactionObject = $(e.relatedTarget).data('object');
                    ServiceTransactions.selectedServiceTransactionData = serviceTransactionObject;
                    self._mapServiceTransactionModal(serviceTransactionObject);
                    self._mapAdditionalCostTableModal(serviceTransactionObject);
                    $('#btnServiceTransactionEditSave').attr('data-mode', 'edit');
                    $('#service-transaction-id').val(serviceTransactionId);
                } else {
                    // Add
                    $('#btnServiceTransactionEditSave').attr('data-mode', 'add');
                    self._mapServiceTransactionModal(null);
                }
            });

            $serviceTransactionEditModal.on('hide.bs.modal', function(e) {
                $('#service-transaction-form').trigger('reset');
            });

            $(document).on('keyup', '#serviceTransactionEditModal #service-transaction-price', function(e) {
                e.preventDefault();
                ServiceTransactions.serviceCost = accounting.unformat($(this).val(), ',');
                self._calculateAllCostTotal();
            });

            $(document).on('click', '#btnServiceTransactionDelete', function(e) {
                e.preventDefault();
                console.log($('#btnServiceTransactionDelete').attr('data-id'));
                let url = base_url + '/api/service-transactions/' + $('#btnServiceTransactionDelete').attr('data-id');

                APIs.ServiceTransactionDelete($('#btnServiceTransactionDelete').attr('data-id'), function(res) {
                    $('#table-service-transactions').DataTable().ajax.reload();
                    $('#serviceTransactionDeleteModal').modal('hide');
                });
            });

            $(document).on('click', '#btnSendStatusWhatsapp', function(e) {
                e.preventDefault();
                let object = JSON.parse($('#customer-object').val());
                let number = $('#customer-whatsapp-number').val().toString();

                let message;
                switch(object.status) {
                    case 'new' :
                        message = `Bapak/Ibu ${object.customer.name},\n saat ini data service anda telah masuk ke sistem kami. Terima kasih.`;
                        break;
                    case 'in-progress':
                        message = `Bapak/Ibu ${object.customer.name},\n kami telah memulai pengerjaan service barang yang telah anda titipkan ke kami. Terima kasih.`;
                        break;
                    case 'completed':
                        message = `Bapak/Ibu ${object.customer.name}, service telah selesai, anda bisa mengambil barang service anda. Terima kasih.`;
                        break;
                    default:
                        message = `Bapak/Ibu ${object.customer.name}, service telah selesai, anda bisa mengambil barang service anda. Terima kasih.`;
                }

                let whatsappUrl = `https://web.whatsapp.com/send?phone=${number}&text=${message}`;
                window.open(whatsappUrl, '_blank');
                window.focus();

                $('#whatsappConfirmationModal').modal('hide');
            });

            $(document).on('click', '#btnServiceTransactionEditSave', function(e) {
                e.preventDefault();

                console.log('prepare payload for edit data');

                // Update selectedServiceTransactionData
                ServiceTransactions.selectedServiceTransactionData.service_date = $('#service-transaction-date').val();
                ServiceTransactions.selectedServiceTransactionData.invoice_no = $('#service-transaction-invoice-no').val();
                ServiceTransactions.selectedServiceTransactionData.item_name = $('#service-transaction-item-name').val();
                ServiceTransactions.selectedServiceTransactionData.damage_type = $('#service-transaction-damage-type').val();
                ServiceTransactions.selectedServiceTransactionData.equipment = $('#service-transaction-equipment').val().join(',');
                ServiceTransactions.selectedServiceTransactionData.description = $('#service-transaction-description').val();
                ServiceTransactions.selectedServiceTransactionData.technician = $('#service-transaction-technician').val();
                ServiceTransactions.selectedServiceTransactionData.repair_type = $('#service-transaction-repair-type').val();
                ServiceTransactions.selectedServiceTransactionData.spare_part = $('#service-transaction-spare-part').val();
                ServiceTransactions.selectedServiceTransactionData.price = accounting.unformat($('#service-transaction-price').val(), ',').toString();
                ServiceTransactions.selectedServiceTransactionData.total_price = accounting.unformat($('#service-transaction-total-price').val(), ",").toString();
                ServiceTransactions.selectedServiceTransactionData.taken_date = $('#service-transaction-taken-date').val();
                ServiceTransactions.selectedServiceTransactionData.status = $('#service-transaction-status').val();
                ServiceTransactions.selectedServiceTransactionData.id = ServiceTransactions.selectedServiceTransactionData.id.toString();
                // ServiceTransactions.selectedServiceTransactionData.customer_id = ServiceTransactions.selectedServiceTransactionData.customer_id.toString();
                ServiceTransactions.selectedServiceTransactionData.customer_id = $('#service-transaction-customer').val().toString();

                delete ServiceTransactions.selectedServiceTransactionData.customer;
                delete ServiceTransactions.selectedServiceTransactionData.additional_items;

                var payload = {
                    serviceTransaction: ServiceTransactions.selectedServiceTransactionData,
                    additionalItems: ServiceTransactions.additionalItemsTableData,
                };

                $.each(payload.additionalItems, function(key, value) {
                    payload.additionalItems[key].id = payload.additionalItems[key].id.toString();
                    payload.additionalItems[key].cost = payload.additionalItems[key].cost.toString();
                    payload.additionalItems[key].st_id = payload.additionalItems[key].st_id.toString();
                });

                APIs.ServiceTransactionUpdate(payload, ServiceTransactions.selectedServiceTransactionData.id, function(res) {
                    $('#table-service-transactions').DataTable().ajax.reload();
                    $('#serviceTransactionEditModal').modal('hide');
                });
            });

            $(document).on('click', '#btnServiceTransactionCreateInvoice', function(e) {
                e.preventDefault();

                console.log('create invoice');
            });

            $(document).on('click', '#btnAdditionalCostSave', function(e) {
                e.preventDefault();

                let mode = $('#btnAdditionalCostSave').attr('data-mode');
                if (mode === 'add') {
                    // add
                    let newAdditionalCost = {
                        "id": uniqueID(),
                        "name": $('#additional-cost-name').val(),
                        "notes": $('#additional-cost-note').val(),
                        "cost": accounting.unformat($('#additional-cost-price').val(),
                            ",").toString(),
                        "st_id": $('#service-transaction-id').val(),
                    };
                    // console.log(newAdditionalCost);
                    ServiceTransactions.additionalItemsTableData.push(newAdditionalCost);
                    // console.log(ServiceTransactions.additionalItemsTableData);
                } else {
                    // edit
                    let key = $('#additional-cost-id').val();
                    let existingIndex = ServiceTransactions.additionalItemsTableData.findIndex(({id}) => id === key);
                    console.log(key, existingIndex);
                    let updatedAdditionalCost = {
                        "id": key,
                        "name": $('#additional-cost-name').val(),
                        "notes": $('#additional-cost-note').val(),
                        "cost": accounting.unformat($('#additional-cost-price').val(), ",").toString(),
                        "st_id": $('#service-transaction-id').val(),
                    };
                    ServiceTransactions.additionalItemsTableData[existingIndex] = updatedAdditionalCost;
                    // console.log(ServiceTransactions.additionalItemsTableData);
                }
                $('#table-additional-cost-modal').DataTable().ajax.reload();
                $('#additional-cost-form').trigger('reset');
            });

            $(document).on('click', '#btnAdditionalCostDelete', function(e) {
                e.preventDefault();
                let key = $('#btnAdditionalCostDelete').attr('data-key');
                ServiceTransactions.additionalItemsTableData.splice(ServiceTransactions.additionalItemsTableData.findIndex(({id}) => id === key.toString()), 1);
                $('#table-additional-cost-modal').DataTable().ajax.reload();
                $('#additionalCostDeleteModal').modal('hide');
            });
        },
        _mapAdditionalCostTableModal: function(serviceTransaction) {
            var self = this;
            // console.log(serviceTransaction);

            if (serviceTransaction !== null) {
                console.log(serviceTransaction.additional_items);
                if (serviceTransaction.additional_items.length > 0) {
                    ServiceTransactions.additionalItemsTableData = [...serviceTransaction.additional_items];
                    $.each(ServiceTransactions.additionalItemsTableData, function(key, value) {
                       ServiceTransactions.additionalItemsTableData[key].id = ServiceTransactions.additionalItemsTableData[key].id.toString();
                    });
                    // console.log(ServiceTransactions.additionalItemsTableData);
                } else {
                    ServiceTransactions.additionalItemsTableData = [];
                }
                $('#table-additional-cost-modal').DataTable().ajax.reload();
            }
        },

        _mapServiceTransactionModal: function(serviceTransaction) {
            var self = this;
            if (serviceTransaction !== null) {
                // $('#serviceTransactionEditModal #service-transaction-date').val(serviceTransaction.service_date);
                $('#serviceTransactionEditModal #service-transaction-date').val(moment(serviceTransaction.service_date).format('DD/MM/YYYY'));
                $('#serviceTransactionEditModal #service-transaction-invoice-no').val(serviceTransaction.invoice_no);
                // $('#serviceTransactionEditModal #service-transaction-customer').val(serviceTransaction.customer_id);
                $('#serviceTransactionEditModal #service-transaction-item-name').val(serviceTransaction.item_name);
                $('#serviceTransactionEditModal #service-transaction-damage-type').val(serviceTransaction.damage_type);
                $('#serviceTransactionEditModal #service-transaction-equipment').val(serviceTransaction.equipment.split(',')).trigger('change');
                $('#serviceTransactionEditModal #service-transaction-description').val(serviceTransaction.description);
                $('#serviceTransactionEditModal #service-transaction-technician').val(serviceTransaction.technician);
                $('#serviceTransactionEditModal #service-transaction-repair-type').val(serviceTransaction.repair_type);
                $('#serviceTransactionEditModal #service-transaction-spare-part').val(serviceTransaction.spare_part);
                $('#serviceTransactionEditModal #service-transaction-price').val(serviceTransaction.price);
                $('#serviceTransactionEditModal #service-transaction-total-price').val(serviceTransaction.total_price);
                $('#serviceTransactionEditModal #service-transaction-taken-date').val(moment(serviceTransaction.taken_date).format('DD/MM/YYYY'));
                $('#serviceTransactionEditModal #service-transaction-status').val(serviceTransaction.status).trigger('change');

                ServiceTransactions.serviceCost = accounting.unformat($('#serviceTransactionEditModal #service-transaction-price').val(), ',');
                self._calculateAllCostTotal();
                self._mapSelectedCustomer(ServiceTransactions.customers, serviceTransaction.customer_id);
            } else {
                $('#service-transaction-form').trigger('reset');
            }
            this._currencyHandler();
        },
        _currencyHandler: function() {
            console.log('format currency on total price input');
            $('#serviceTransactionEditModal #service-transaction-price').priceFormat({
                prefix: '',
                centsSeparator: ',',
                thousandsSeparator: '.',
                clearOnEmpty: true,
                centsLimit: 0,
            });

            $('#serviceTransactionEditModal #service-transaction-total-price').priceFormat({
                prefix: '',
                centsSeparator: ',',
                thousandsSeparator: '.',
                clearOnEmpty: true,
                centsLimit: 0,
            });

            $('#additional-cost-price').priceFormat({
                prefix: '',
                centsSeparator: ',',
                thousandsSeparator: '.',
                clearOnEmpty: true,
                centsLimit: 0,
            });
        },
        _select2Handler: function() {
            $('#serviceTransactionEditModal #service-transaction-status').select2();

            $('#serviceTransactionEditModal #service-transaction-equipment').select2({
                tags: true,
            });
        },
        _modalHandler: function() {
            AddServiceTransactions._modalHandler();
        },
        init: function() {
            this._datatables();
            this._APIs();
            this._handleButtonEvents();
            this._datepicker();
            this._select2Handler();
            this._modalHandler();
        }
    };
});