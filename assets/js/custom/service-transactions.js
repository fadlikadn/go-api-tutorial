"use strict";

$(function() {
    ServiceTransactions = {
        _datatables: function() {
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
                            // TODO parse date and time using MomentJS
                            return data;
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
                    {
                        "data": "price", "defaultContent": ""
                    },
                    {
                        "data": "total_price", "defaultContent": ""
                    },
                    {
                        "data": "status", "defaultContent": ""
                    },
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#serviceTransactionEditModal' data-toggle="modal" class='btn btn-sm btn-success service-transaction_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#serviceTransactionDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger service-transaction_delete' data-key=${data} >Delete</a>`;
                        }
                    }
                ]
            });
        },
        _APIs: function() {

        },
        _handleButtonEvents: function() {
            // Service Transaction Handle Button Events
            // TODO implement button handler for service transaction management page
            var self = this;
            var $serviceTransactionEditModal = $('#serviceTransactionEditModal');
            var $serviceTransactionDeleteModal = $('#serviceTransactionDeleteModal');

            $serviceTransactionDeleteModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Delete
                    var serviceTransactionId = $(e.relatedTarget).data('key');
                    console.log(serviceTransactionId);
                    $('#btnServiceTransactionDelete').attr('data-id', serviceTransactionId);
                }
            });

            $serviceTransactionEditModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Edit
                    var serviceTransactionId = $(e.relatedTarget).data('key');
                    var serviceTransactionObject = $(e.relatedTarget).data('object');
                    console.log(serviceTransactionId, serviceTransactionObject);
                    console.log('prepare for edit');
                    self._mapServiceTransactionModal(serviceTransactionObject);
                    $('#btnServiceTransactionEditSave').attr('data-mode', 'edit');
                    $('#service-transaction-id').val(serviceTransactionId);
                } else {
                    // ADd
                    console.log('prepare for add');
                    $('#btnServiceTransactionEditSave').attr('data-mode', 'add');
                    self._mapServiceTransactionModal(null);
                }
            });

            $serviceTransactionEditModal.on('hide.bs.modal', function(e) {
                $('#service-transaction-form').trigger('reset');
            });
        },
        _mapServiceTransactionModal: function(serviceTransaction) {
            // TODO implement modal for service transaction
            if (serviceTransaction !== null) {
                $('#serviceTransactionEditModal #service-transaction-date').val(serviceTransaction.service_date);
                $('#serviceTransactionEditModal #service-transaction-invoice-no').val(serviceTransaction.invoice_no);
                $('#serviceTransactionEditModal #service-transaction-customer').val(serviceTransaction.customer_id);
                $('#serviceTransactionEditModal #service-transaction-item-name').val(serviceTransaction.item_name);
                $('#serviceTransactionEditModal #service-transaction-damage-type').val(serviceTransaction.damage_type);
                $('#serviceTransactionEditModal #service-transaction-equipment').val(serviceTransaction.equipment);
                $('#serviceTransactionEditModal #service-transaction-description').val(serviceTransaction.description);
                $('#serviceTransactionEditModal #service-transaction-technician').val(serviceTransaction.technician);
                $('#serviceTransactionEditModal #service-transaction-repair-type').val(serviceTransaction.repair_type);
                $('#serviceTransactionEditModal #service-transaction-price').val(serviceTransaction.price);
                $('#serviceTransactionEditModal #service-transaction-total-price').val(serviceTransaction.total_price);
                $('#serviceTransactionEditModal #service-transaction-taken-date').val(serviceTransaction.taken_date);
                $('#serviceTransactionEditModal #service-transaction-status').val(serviceTransaction.status);
            } else {
                $('#service-transaction-form').trigger('reset');
            }
        },
        init: function() {
            this._datatables();
            this._APIs();
            this._handleButtonEvents();
        }
    };

    AddServiceTransactions = {
        customerList: null,
        additionalItems: [
            {
                "id": '6712871827',
                "name": "Mainboard",
                "notes": "Ganti Keyboard merk Acer",
                "cost": 600000,
            }
        ],
        additionalCostTable: null,
        additionalCostTotal: 0,
        serviceCost: 0,
        allCostTotal: 0,
        _wizardHandler: function() {

            // Toolbar extra buttons
            var btnFinish = $('<button></button>').text('Finish')
                .addClass('btn btn-info')
                .on('click', function(){ alert('Finish Clicked'); });
            var btnCancel = $('<button></button>').text('Cancel')
                .addClass('btn btn-danger')
                .on('click', function(){ $('#smartwizard').smartWizard("reset"); });

            $('#service-transaction-smartwizard').smartWizard({
                selected: 0,
                keyNavigation:false,
                theme: 'arrows',
                transitionEffect: 'fade',
                showStepURLhash: false,
                toolbarSettings: {
                    // toolbarPosition: 'both',
                    toolbarButtonPosition: 'end',
                    // toolbarExtraButtons: [btnFinish, btnCancel]
                },
                lang: {
                    next: 'Lanjut',
                    previous: 'Kembali',
                }
            });

            $('#service-transaction-smartwizard').on('leaveStep', function(e, anchorObject, stepNumber, stepDirection) {
                if (stepNumber === 0) {
                    let customerMode = $('#customer-mode').val();
                    console.log($('#customer-mode').val());
                    if (stepDirection === 'forward') {
                        if (customerMode === 'search-existing') {
                            let selectedCustomer = $('#existing-customer :selected').val();
                            console.log(selectedCustomer);
                            if (selectedCustomer === '') {
                                alert('customer belum dipilih');
                                return false;
                            }
                        } else if (customerMode === 'add-new') {
                            let customerForm = $('#customer-form');
                            customerForm.validator('validate');
                            var elmErr = customerForm.children('.has-error');
                            if (elmErr && elmErr.length > 0) {
                                alert('form belum lengkap, mohon diisi');
                                return false;
                            }
                        }
                    }

                    // if (stepDirection === 'forward' && elmForm) {
                    //     elmForm.validator('validate');
                    //     var elmErr = elmForm.children('.has-error');
                    //     if (elmErr && elmErr.length > 0) {
                    //         // Form validation failed
                    //         return false;
                    //     }
                    // }
                }
                if (stepNumber === 1) {

                }
            });
        },
        _datepicker: function() {
            console.log('setup datepicker');
            $('#service-transaction-date').datetimepicker({
                format: 'DD/MM/YYYY',
            });

            $('#service-transaction-taken-date').datetimepicker({
                format: 'DD/MM/YYYY'
            });
        },
        _select2Handler: function() {
            $('#customer-mode').select2();

            $('#service-transaction-equipment').select2({
                tags: true,
            });

            $('#existing-customer-container').removeClass('d-none');

            $('#customer-mode').on('select2:select', function(e) {
                var data = e.params.data;

                if (data.id === 'add-new') {
                    $('#new-customer-container').removeClass('d-none');
                    $('#existing-customer-container').addClass('d-none');

                    $('#customer-existing-alert').addClass('d-none');
                    $('#customer-existing-found').empty();

                    $('#existing-customer').val(null).trigger('change.select2');
                } else if (data.id === 'search-existing') {
                    $('#existing-customer-container').removeClass('d-none');
                    $('#new-customer-container').addClass('d-none');

                    $('#customer-form').trigger('reset');
                }
                console.log(data);
            });
        },
        _currencyHandler: function() {
            $('#service-transaction-price').priceFormat({
                prefix: '',
                centsSeparator: ',',
                thousandsSeparator: '.',
                clearOnEmpty: true,
            });

            $('#additional-cost-price').priceFormat({
                prefix: '',
                centsSeparator: ',',
                thousandsSeparator: '.',
                clearOnEmpty: true
            });
        },
        _APIs: function() {
            var self = this;
            $(document).on('click', '#btnCustomerRefresh', function() {
                console.log('btnCustomRefresh clicked');
                APIs.CustomerGetAll(function(res) {
                    Customers.customerList = res;
                    console.log(Customers.customerList);
                    self._mapCustomerList(Customers.customerList);

                    $('#customer-existing-alert').addClass('d-none');
                });
            });
            $('#btnCustomerRefresh').trigger('click');
        },
        _mapCustomerList: function(customers) {
            var $existingCustomer = $('#existing-customer');
            $existingCustomer.empty();
            $existingCustomer.append($("<option/>"));
            $.each(customers, function(key, value) {
                var $option = $("<option/>", {
                    value: value.id,
                    text: value.name,
                    selected: false,
                    object: JSON.stringify(value),
                });
                $existingCustomer.append($option);
            });
            $existingCustomer.select2({
                placeholder: "Select Customer",
                allowClear: true
            });

            $existingCustomer.on('select2:select', function(e) {
                // var data = e.params.data;
                let object = $(e.params.data.element).attr('object');
                object = JSON.parse(object);
                console.log(object);

                // fill element customer-existing-found
                $('#customer-existing-alert').removeClass('d-none');
                $('#customer-existing-found').html(`Nama: ${object.name}, Email: ${object.email !== '' ? object.email : '-'}, Telepon: ${object.phone !== '' ? object.phone : '-'}`);

            });

            $existingCustomer.on('select2:clear', function(e) {
                $('#customer-existing-alert').addClass('d-none');
                $('#customer-existing-found').empty();
            });
        },
        _datatables: function() {
            var self = this;
            AddServiceTransactions.additionalCostTable =  $('#table-additional-cost').on('preXhr.dt', function(e, settings, data) {
                console.log('reset additionalCostTotal before table load new data');

                AddServiceTransactions.additionalCostTotal = 0;
                AddServiceTransactions.allCostTotal = 0;
                AddServiceTransactions.allCostTotal += AddServiceTransactions.serviceCost;
                self._calculateAllCostTotal();
            })
                .DataTable({
                "processing": true,
                "searching": true,
                "serverSide": false,
                "paging": true,
                "bLengthChange": true,
                "ordering": true,
                // "data": AddServiceTransactions.additionalItems,
                "ajax": function(data, callback, settings) {
                    callback({ data: AddServiceTransactions.additionalItems})
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
                            return accounting.formatMoney(data, "Rp", 0, ".", ",");
                        }
                    },
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#additionalCostModal' data-toggle="modal" class='btn btn-sm btn-success additional-cost_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#additionalCostDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger additional-cost_delete' data-key=${data} >Delete</a>`;
                        }
                    }
                ],
                // "preInit": function() {
                //     AddServiceTransactions.additionalCostTotal = 0;
                // },
                "rowCallback": function(row, data, index) {
                    AddServiceTransactions.additionalCostTotal += data.cost;
                    console.log(AddServiceTransactions.additionalCostTotal);
                    self._calculateAllCostTotal();
                },
            });
        },
        _calculateAllCostTotal: function() {
            AddServiceTransactions.allCostTotal = AddServiceTransactions.serviceCost + AddServiceTransactions.additionalCostTotal;
            $('#total-cost').html(accounting.formatMoney(AddServiceTransactions.allCostTotal, "Rp", 0, ".", ","));
        },
        _buttonsHandler: function() {
            var self = this;
            $(document).on('click', '#btnAdditionalCostSave', function(e) {
                e.preventDefault();

                let mode = $('#btnAdditionalCostSave').attr('data-mode');
                if (mode === 'add') {
                    // add
                    // Store new Additional Item
                    let newAdditionalCost = {
                        "id": uniqueID(),
                        "name": $('#additional-cost-name').val(),
                        "notes": $('#additional-cost-note').val(),
                        "cost": accounting.unformat($('#additional-cost-price').val(), ","),
                    };
                    console.log(newAdditionalCost);
                    AddServiceTransactions.additionalItems.push(newAdditionalCost);
                    console.log(AddServiceTransactions.additionalItems);
                    $('#table-additional-cost').DataTable().ajax.reload();
                    $('#additional-cost-form').trigger('reset');
                } else {
                    // edit
                    let key = $('#additional-cost-id').val();
                    let existingIndex = AddServiceTransactions.additionalItems.findIndex(({id}) => id === key);
                    console.log(key, existingIndex);
                    let updatedAdditionalCost = {
                        "id": key,
                        "name": $('#additional-cost-name').val(),
                        "notes": $('#additional-cost-note').val(),
                        "cost": accounting.unformat($('#additional-cost-price').val(), ","),
                    };
                    AddServiceTransactions.additionalItems[existingIndex] = updatedAdditionalCost;
                    console.log(AddServiceTransactions.additionalItems[existingIndex]);
                    console.log(AddServiceTransactions.additionalItems);
                    $('#table-additional-cost').DataTable().ajax.reload();
                    $('#additional-cost-form').trigger('reset');
                }

            });

            $(document).on('click', '#btnAdditionalCostDelete', function(e) {
                e.preventDefault();
                console.log($('#btnAdditionalCostDelete').attr('data-key'));
                var key = $('#btnAdditionalCostDelete').attr('data-key');
                AddServiceTransactions.additionalItems.splice(AddServiceTransactions.additionalItems.findIndex(({id}) => id === key), 1);
                $('#table-additional-cost').DataTable().ajax.reload();
                $('#additionalCostDeleteModal').modal('hide');
            });

            $(document).on('keyup', '#service-transaction-price', function(e) {
                e.preventDefault();
                console.log(accounting.unformat($(this).val(), ','));
                AddServiceTransactions.serviceCost = accounting.unformat($(this).val(), ',');
                self._calculateAllCostTotal();
            });
        },
        _modalHandler: function() {
            var self = this;
            var $additionalCostModal = $('#additionalCostModal');

            $additionalCostModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') !== undefined) {
                    // Edit
                    let additionalCostId = $(e.relatedTarget).data('key');
                    let additionalCostObject = $(e.relatedTarget).data('object');
                    console.log(additionalCostId, additionalCostObject);
                    $('#btnAdditionalCostSave').attr('data-mode', 'edit');
                    $('#additional-cost-id').val(additionalCostId);
                    self._mapAdditionalCostModal(additionalCostObject);
                } else {
                    // Add
                    $('#btnAdditionalCostSave').attr('data-mode', 'add');
                    self._mapAdditionalCostModal(null);
                }
            });
        },
        _mapAdditionalCostModal: function(additionalCost) {
            if (additionalCost !== null) {
                $('#additional-cost-form #additional-cost-name').val(additionalCost.name);
                $('#additional-cost-form #additional-cost-note').val(additionalCost.notes);
                $('#additional-cost-form #additional-cost-price').val(accounting.formatMoney(additionalCost.cost, '', 2, '.', ','));
            } else {
                $('#additional-cost-form').trigger('reset');
            }
        },
        init: function() {
            this._wizardHandler();
            this._select2Handler();
            this._APIs();
            this._datepicker();
            this._currencyHandler();
            this._datatables();
            this._buttonsHandler();
            this._modalHandler();
        }
    };
});