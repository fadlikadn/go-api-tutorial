"use strict";

$(function() {
    AddServiceTransactions = {
        customerList: null,
        additionalItems: [
            // {
            //     "id": '6712871827',
            //     "name": "Mainboard",
            //     "notes": "Ganti Keyboard merk Acer",
            //     "cost": 600000,
            // }
        ],
        additionalCostTable: null,
        additionalCostTotal: 0,
        serviceCost: 0,
        allCostTotal: 0,
        customerDataObject: {},
        serviceTransactionDataObject: {},
        _wizardHandler: function() {
            var self = this;
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
                autoAdjustHeight: false,
                showStepURLhash: false,
                toolbarSettings: {
                    toolbarPosition: 'both',
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
                                alert('mohon lengkapi data customer');
                                return false;
                            }
                        }
                    }
                }
                if (stepNumber === 1) {
                    // Validate Service Transaction Data
                    let serviceTransactionForm = $('#service-transaction-form');
                    serviceTransactionForm.validator('validate');
                    var elmErr = serviceTransactionForm.children('.has-error');
                    if (elmErr && elmErr.length > 0) {
                        alert('mohon lengkapi data service transction');
                        return false;
                    }

                    // Process data customer and service-transaction into object
                    self._previewCustomer();
                    self._previewServiceTransaction();
                    self._previewAdditionalItem();
                }
            });
        },
        _previewAdditionalItem: function() {
            let $previewAdditionalItemList = $('#preview-additional-item-list');
            $previewAdditionalItemList.empty();
            $.each(AddServiceTransactions.additionalItems, function(key, value) {
                var $li = `<li>${value.name} : ${accounting.formatMoney(value.cost, 'Rp ', 0, '.', ',')}</li>`;
                $previewAdditionalItemList.append($li);
            });
        },
        _previewServiceTransaction: function() {
            AddServiceTransactions.serviceTransactionDataObject.service_date = $('#service-transaction-date').val();
            AddServiceTransactions.serviceTransactionDataObject.invoice_no = $('#service-transaction-invoice-no').val();
            AddServiceTransactions.serviceTransactionDataObject.item_name = $('#service-transaction-item-name').val();
            AddServiceTransactions.serviceTransactionDataObject.damage_type = $('#service-transaction-damage-type').val();
            AddServiceTransactions.serviceTransactionDataObject.equipment = $('#service-transaction-equipment').val();
            AddServiceTransactions.serviceTransactionDataObject.description = $('#service-transaction-description').val();
            AddServiceTransactions.serviceTransactionDataObject.technician = $('#service-transaction-technician').val();
            AddServiceTransactions.serviceTransactionDataObject.repair_type = $('#service-transaction-repair-type').val();
            AddServiceTransactions.serviceTransactionDataObject.spare_part = $('#service-transaction-spare-part').val();
            AddServiceTransactions.serviceTransactionDataObject.price = accounting.unformat($('#service-transaction-price').val(), ',');
            AddServiceTransactions.serviceTransactionDataObject.total_price = accounting.unformat($('#total-cost').html(), ",");
            AddServiceTransactions.serviceTransactionDataObject.taken_date = $('#service-transaction-taken-date').val();
            AddServiceTransactions.serviceTransactionDataObject.status = 'new';

            console.log(AddServiceTransactions.serviceTransactionDataObject);

            $('#preview-service-transaction-date').html(AddServiceTransactions.serviceTransactionDataObject.service_date);
            $('#preview-service-transaction-invoice-no').html(AddServiceTransactions.serviceTransactionDataObject.invoice_no);
            $('#preview-service-transaction-item-name').html(AddServiceTransactions.serviceTransactionDataObject.item_name);
            $('#preview-service-transaction-damage-type').html(AddServiceTransactions.serviceTransactionDataObject.damage_type);
            $('#preview-service-transaction-equipment').html(AddServiceTransactions.serviceTransactionDataObject.equipment);
            $('#preview-service-transaction-description').html(AddServiceTransactions.serviceTransactionDataObject.description);
            $('#preview-service-transaction-technician').html(AddServiceTransactions.serviceTransactionDataObject.technician);
            $('#preview-service-transaction-repair-type').html(AddServiceTransactions.serviceTransactionDataObject.repair_type);
            $('#preview-service-transaction-spare-part').html(AddServiceTransactions.serviceTransactionDataObject.spare_part);
            $('#preview-service-transaction-price').html(accounting.formatMoney(AddServiceTransactions.serviceTransactionDataObject.price, "Rp ", 0, '.', ','));
            $('#preview-service-transaction-taken-date').html(AddServiceTransactions.serviceTransactionDataObject.taken_date);
            $('#preview-service-transaction-status').html(AddServiceTransactions.serviceTransactionDataObject.status);
            $('#preview-total-cost').html(accounting.formatMoney(AddServiceTransactions.serviceTransactionDataObject.total_price, "Rp ", 0, '.', ','));
        },
        _previewCustomer: function() {
            let customerExistingId = $('#customer-existing-id').val();
            let $customerExistingObject = $('#customer-existing-object');
            let customerExistingObject = ($customerExistingObject.val() !== '') ?  JSON.parse($customerExistingObject.val()) : null;
            console.log(customerExistingObject);
            if (customerExistingId !== '') {
                AddServiceTransactions.customerDataObject = customerExistingObject;
            } else {
                // New Customer
                AddServiceTransactions.customerDataObject = {};
                AddServiceTransactions.customerDataObject.id = null;
                AddServiceTransactions.customerDataObject.name = $('#customer-name').val();
                AddServiceTransactions.customerDataObject.email = $('#customer-email').val();
                AddServiceTransactions.customerDataObject.phone = $('#customer-phone').val();
                AddServiceTransactions.customerDataObject.address = $('#customer-address').val();
                AddServiceTransactions.customerDataObject.notes = $('#customer-notes').val();

            }

            console.log(AddServiceTransactions.customerDataObject);

            $('#preview-customer-name').html(AddServiceTransactions.customerDataObject.name);
            $('#preview-customer-email').html(AddServiceTransactions.customerDataObject.email);
            $('#preview-customer-phone').html(AddServiceTransactions.customerDataObject.phone);
            $('#preview-customer-address').html(AddServiceTransactions.customerDataObject.address);
            $('#preview-customer-note').html(AddServiceTransactions.customerDataObject.notes);
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
                    $('#customer-existing-id').val('');
                    $('#customer-existing-object').val('');

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
        _APIs: function() {
            var self = this;
            $(document).on('click', '#btnCustomerRefresh', function() {
                console.log('btnCustomRefresh clicked');
                APIs.CustomerGetAll(function(res) {
                    AddServiceTransactions.customerList = res;
                    console.log(AddServiceTransactions.customerList);
                    self._mapCustomerList(AddServiceTransactions.customerList);

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
                $('#customer-existing-id').val(object.id);
                $('#customer-existing-object').val(JSON.stringify(object));

            });

            $existingCustomer.on('select2:clear', function(e) {
                $('#customer-existing-alert').addClass('d-none');
                $('#customer-existing-found').empty();
                $('#customer-existing-id').val('');
                $('#customer-existing-object').val('');
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
                                return `<span class="float-right">${accounting.formatMoney(data, "Rp", 0, ".", ",")}</span>`;
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
                        AddServiceTransactions.additionalCostTotal += parseInt(data.cost);
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
                        "cost": accounting.unformat($('#additional-cost-price').val(), ",").toString(),
                    };
                    console.log(newAdditionalCost);
                    AddServiceTransactions.additionalItems.push(newAdditionalCost);
                    console.log(AddServiceTransactions.additionalItems);
                } else {
                    // edit
                    let key = $('#additional-cost-id').val();
                    let existingIndex = AddServiceTransactions.additionalItems.findIndex(({id}) => id === key);
                    console.log(key, existingIndex);
                    let updatedAdditionalCost = {
                        "id": key,
                        "name": $('#additional-cost-name').val(),
                        "notes": $('#additional-cost-note').val(),
                        "cost": accounting.unformat($('#additional-cost-price').val(), ",").toString(),
                    };
                    AddServiceTransactions.additionalItems[existingIndex] = updatedAdditionalCost;
                    console.log(AddServiceTransactions.additionalItems[existingIndex]);
                    console.log(AddServiceTransactions.additionalItems);
                }
                $('#table-additional-cost').DataTable().ajax.reload();
                $('#additional-cost-form').trigger('reset');
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

            $(document).on('click', '#btnProcessTransactionAction', function(e) {
                e.preventDefault();
                console.log('process transaction');

                AddServiceTransactions.serviceTransactionDataObject.equipment = AddServiceTransactions.serviceTransactionDataObject.equipment.join(',');
                AddServiceTransactions.serviceTransactionDataObject.price = AddServiceTransactions.serviceTransactionDataObject.price.toString();
                AddServiceTransactions.serviceTransactionDataObject.total_price = AddServiceTransactions.serviceTransactionDataObject.total_price.toString();

                APIs.StoreNewServiceTransaction(
                    AddServiceTransactions.customerDataObject,
                    AddServiceTransactions.serviceTransactionDataObject,
                    AddServiceTransactions.additionalItems, function(res) {
                        console.log(res);
                        $('#confirmationProcessModal').modal('hide');
                        // TODO implement cetak nota/kwitansi (printable / PDF)
                        window.location.replace(base_url + "/dashboard/service-transactions");
                    });

                /*// hit to API
                var payload = {
                    customer: AddServiceTransactions.customerDataObject,
                    serviceTransaction: AddServiceTransactions.serviceTransactionDataObject,
                    additionalItems: AddServiceTransactions.additionalItems,
                };

                console.log(payload);
                console.log(JSON.stringify(payload));*/
            })
        },
        _modalHandler: function() {
            var self = this;
            var $additionalCostModal = $('#additionalCostModal');
            var $additionalCostDeleteModel = $('#additionalCostDeleteModal');

            $additionalCostModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') !== undefined) {
                    // Edit
                    let additionalCostId = $(e.relatedTarget).data('key');
                    let additionalCostObject = $(e.relatedTarget).data('object');
                    console.log(additionalCostId, additionalCostObject);
                    $('#btnAdditionalCostSave').attr('data-mode', 'edit');
                    $('#additional-cost-id').val(additionalCostId);
                    AddServiceTransactions._mapAdditionalCostModal(additionalCostObject);
                } else {
                    // Add
                    $('#btnAdditionalCostSave').attr('data-mode', 'add');
                    AddServiceTransactions._mapAdditionalCostModal(null);
                }
            });

            $additionalCostDeleteModel.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') !== undefined) {
                    // Delete
                    let additionalCostId = $(e.relatedTarget).data('key');
                    console.log(additionalCostId);
                    $('#btnAdditionalCostDelete').attr('data-key', additionalCostId);
                }
            });
        },
        _mapAdditionalCostModal: function(additionalCost) {
            if (additionalCost !== null) {
                $('#additional-cost-form #additional-cost-name').val(additionalCost.name);
                $('#additional-cost-form #additional-cost-note').val(additionalCost.notes);
                $('#additional-cost-form #additional-cost-price').val(accounting.formatMoney(additionalCost.cost, '', 0, '.', ','));
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