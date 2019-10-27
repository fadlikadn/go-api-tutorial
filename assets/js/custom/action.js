"use strict";

var Login = {};
var Register = {};
var Users = {};
var Customers = {};
var ServiceTransactions = {};
var AddServiceTransactions = {};
// TODO Implement JS handler for ServiceTransactions

$(function() {


    Register = {
        _prepareAction: function() {
            $(document).on('click', '#btn_register_action', function(e) {
                e.preventDefault();

                var $self = $(this);
                var payload = JSON.stringify({
                    name: $('#name').val(),
                    email: $('#email').val(),
                    phone: $('#phone').val(),
                    company: $('#company').val(),
                    password: $('#inputPassword').val(),
                });

                $.ajax({
                    url: base_url + '/api/register',
                    method: 'POST',
                    data: payload,
                    contentType: 'application/json'
                }).then(function(res) {
                    console.log(res);
                    console.log('will send you an activation email');
                    window.location.replace(base_url + "/activation-pending");
                }).catch(function(err) {
                    console.log(err);
                    console.log('registration failed, please try again');
                })
            });
        },
        init: function() {
            this._prepareAction();
        }
    };

    Users = {
        _datatables: function() {
            console.log(base_url);
            $('#table-users').DataTable({
                "processing": true,
                "searching": true,
                "serverSide": false,
                "paging": true,
                "bLengthChange": true,
                "ordering": true,
                "ajax": {
                    "url": base_url + "/api/session/users",
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
                        "data": "name", "defaultContent": ""
                    },
                    {
                        "data": "email", "defaultContent": ""
                    },
                    {
                        "data": "phone", "defaultContent": ""
                    },
                    {
                        "data": "company", "defaultContent": ""
                    },
                    {
                        "data": "is_active", "defaultContent": ""
                    },
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#userEditModal' data-toggle="modal" class='btn btn-sm btn-success user_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#userDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger user_delete' data-key=${data} >Delete</a>`;
                        }
                    }
                ]
            });
        },
        _APIs: function() {

        },
        _handleButtonEvents: function() {
            var self = this;
            var $userEditModal = $('#userEditModal');
            var $userDeleteModal = $('#userDeleteModal');

            $userDeleteModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Delete
                    var userId = $(e.relatedTarget).data('key');
                    console.log(userId);
                    $('#btnUserDelete').attr('data-id', userId);
                }
            });

            $userEditModal.on('show.bs.modal', function(e){
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Edit
                    var userId = $(e.relatedTarget).data('key');
                    var userObject = $(e.relatedTarget).data('object');
                    console.log(userId, userObject);
                    console.log('prepare for edit');
                    self._mapUserModal(userObject);
                    $('#btnUserEditSave').attr('data-mode', 'edit');
                    $('#user-id').val(userId);
                } else {
                    // Add
                    console.log('prepare for add');
                    $('#btnUserEditSave').attr('data-mode', 'add');
                    self._mapUserModal(null);
                }
            });

            $userEditModal.on('hide.bs.modal', function(e) {
                $('#user-form').trigger('reset');
            });

            $(document).on('click', '#btnUserDelete', function(e) {
                e.preventDefault();
                console.log($('#btnUserDelete').attr('data-id'));
                var url = base_url + '/api/session/users/' + $('#btnUserDelete').attr('data-id');

                $.ajax({
                    url: url,
                    method: 'DELETE',
                    contentType: 'application/json'
                }).then(function(res) {
                    console.log(res);
                    $('#table-users').DataTable().ajax.reload();
                    $('#userDeleteModal').modal('hide');
                }).catch(function(err) {
                    console.log(err);
                });
            });

            $(document).on('click', '#btnUserEditSave', function(e) {
                e.preventDefault();
                var url = null;
                var payload = null;
                var method = '';
                var mode = $('#btnUserEditSave').attr('data-mode');
                if (mode === 'add') {
                    // add
                    url = base_url + '/api/session/users';
                    method = 'POST';
                    payload = {
                        name: $('#userEditModal #user-name').val(),
                        email: $('#userEditModal #user-email').val(),
                        phone: $('#userEditModal #user-phone').val(),
                        company: $('#userEditModal #user-company').val(),
                        is_active: (parseInt($('#userEditModal #user-is-active').val()) === 1),
                        password: $('#userEditModal #user-password').val(),
                    };
                } else {
                    // edit
                    method = 'PUT';
                    url = base_url + '/api/session/users/' + $('#user-id').val();
                    payload = {
                        name: $('#userEditModal #user-name').val(),
                        email: $('#userEditModal #user-email').val(),
                        phone: $('#userEditModal #user-phone').val(),
                        company: $('#userEditModal #user-company').val(),
                        is_active: (parseInt($('#userEditModal #user-is-active').val()) === 1),
                        password: $('#userEditModal #user-password').val(),
                    };
                }

                $.ajax({
                    url: url,
                    method: method,
                    data: JSON.stringify(payload),
                    contentType: 'application/json'
                }).then(function(res) {
                    console.log(res);
                    $('#table-users').DataTable().ajax.reload();
                    $('#userEditModal').modal('hide');
                }).catch(function(err) {
                    console.log(err);
                });
            });
        },
        _mapUserModal: function(user) {
            if (user !== null) {
                $('#userEditModal #user-name').val(user.name);
                $('#userEditModal #user-email').val(user.email);
                $('#userEditModal #user-phone').val(user.phone);
                $('#userEditModal #user-company').val(user.company);
                $('#userEditModal #user-is-active').val((user.is_active) ? 1 : 0);
            } else {
                $('#user-form').trigger('reset');
            }

        },
        init: function() {
            this._datatables();
            this._APIs();
            this._handleButtonEvents();
        }
    };

    Login = {
        _prepareAction: function() {
            $(document).on('click', '#btn_login_action', function(e) {
                e.preventDefault();

                var $self = $(this);
                var payload = JSON.stringify({
                    email: $('#inputEmail').val(),
                    password: $('#inputPassword').val()
                });

                $.ajax({
                    url: base_url + '/api/login',
                    method: 'POST',
                    data: payload,
                    contentType: 'application/json'
                }).then(function(res) {
                    console.log(res);
                    // alert('berhasil \n' + res);
                    window.location.replace(base_url+'/');
                }).catch(function(err) {
                    console.log(err);
                    alert('gagal');
                });
            });
        },
        init: function() {
            this._prepareAction();
        },
    };

    Customers = {
        _datatables: function() {
            $('#table-customers').DataTable({
                "processing": true,
                "searching": true,
                "serverSide": false,
                "paging": true,
                "bLengthChange": true,
                "ordering": true,
                "ajax": {
                    "url": base_url + "/api/customers",
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
                        "data": "name", "defaultContent": ""
                    },
                    {
                        "data": "email", "defaultContent": ""
                    },
                    {
                        "data": "phone", "defaultContent": ""
                    },
                    {
                        "data": "address", "defaultContent": ""
                    },
                    {
                        "data": "notes", "defaultContent": ""
                    },
                    {
                        "data": "id", render: function(data, type, row, meta) {
                            return `<a href='#' data-target='#customerEditModal' data-toggle="modal" class='btn btn-sm btn-success customer_edit' data-key=${data} data-object='${JSON.stringify(row)}'>Edit</a> &nbsp; <a href='#' data-target='#customerDeleteModal' data-toggle="modal" class='btn btn-sm btn-danger customer_delete' data-key=${data} >Delete</a>`;
                        }
                    }
                ]
            });
        },
        _APIs: function() {

        },
        _handleButtonEvents: function() {
            var self = this;
            var $customerEditModal = $('#customerEditModal');
            var $customerDeleteModal = $('#customerDeleteModal');

            $customerDeleteModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Delete
                    var customerId = $(e.relatedTarget).data('key');
                    console.log(customerId);
                    $('#btnCustomerDelete').attr('data-id', customerId);
                }
            });

            $customerEditModal.on('show.bs.modal', function(e) {
                if ($(e.relatedTarget).data('key') != undefined) {
                    // Edit
                    var customerId = $(e.relatedTarget).data('key');
                    var customerObject = $(e.relatedTarget).data('object');
                    console.log(customerId, customerObject);
                    console.log('prepare for edit');
                    self._mapCustomerModal(customerObject);
                    $('#btnCustomerEditSave').attr('data-mode', 'edit');
                    $('#customer-id').val(customerId);
                } else {
                    // Add
                    console.log('prepare for add');
                    $('#btnCustomerEditSave').attr('data-mode', 'add');
                    self._mapCustomerModal(null);
                }
            });

            $customerEditModal.on('hide.bs.modal', function(e) {
                $('#customer-form').trigger('reset');
            });

            $(document).on('click', '#btnCustomerDelete', function(e) {
                e.preventDefault();
                console.log($('#btnCustomerDelete').attr('data-id'));
                var url = base_url + '/api/customers/' + $('#btnCustomerDelete').attr('data-id');

                $.ajax({
                    url: url,
                    method: 'DELETE',
                    contentType: 'application/json',
                }).then(function(res) {
                    console.log(res);
                    $('#table-customers').DataTable().ajax.reload();
                    $('#customerDeleteModal').modal('hide');
                }).catch(function(err) {
                    console.log(err);
                });
            });

            $(document).on('click', '#btnCustomerEditSave', function(e) {
                e.preventDefault();
                var url = null;
                var payload = null;
                var method = '';
                var mode = $('#btnCustomerEditSave').attr('data-mode');
                if (mode === 'add') {
                    // add
                     url = base_url + '/api/customers';
                     method = 'POST';
                     payload = {
                         name: $('#customerEditModal #customer-name').val(),
                         email: $('#customerEditModal #customer-email').val(),
                         phone: $('#customerEditModal #customer-phone').val(),
                         address: $('#customerEditModal #customer-address').val(),
                         notes: $('#customerEditModal #customer-notes').val(),
                     };
                } else {
                    // edit
                    url = base_url + '/api/customers/' + $('#customer-id').val();
                    method = 'PUT';
                    payload = {
                        name: $('#customerEditModal #customer-name').val(),
                        email: $('#customerEditModal #customer-email').val(),
                        phone: $('#customerEditModal #customer-phone').val(),
                        address: $('#customerEditModal #customer-address').val(),
                        notes: $('#customerEditModal #customer-notes').val(),
                    };
                }

                $.ajax({
                    url: url,
                    method: method,
                    data: JSON.stringify(payload),
                    contentType: 'application/json'
                }).then(function(res) {
                    console.log(res);
                    $('#table-customers').DataTable().ajax.reload();
                    $('#customerEditModal').modal('hide');
                }).catch(function(err) {
                    console.log(err);
                });
            });
        },
        _mapCustomerModal: function(customer) {
            if (customer !== null) {
                $('#customerEditModal #customer-name').val(customer.name);
                $('#customerEditModal #customer-email').val(customer.email);
                $('#customerEditModal #customer-phone').val(customer.phone);
                $('#customerEditModal #customer-address').val(customer.address);
                $('#customerEditModal #customer-notes').val(customer.notes);
            } else {
                $('#customer-form').trigger('reset');
            }
        },
        init: function() {
            this._datatables();
            this._APIs();
            this._handleButtonEvents();
        }
    };

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
                theme: 'arrows',
                transitionEffect: 'fade',
                showStepURLhash: false,
                toolbarSettings: {
                    // toolbarPosition: 'both',
                    toolbarButtonPosition: 'end',
                    // toolbarExtraButtons: [btnFinish, btnCancel]
                },
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
        _select2Handler: function() {
            $('#customer-mode').select2();

            $('#existing-customer-container').removeClass('d-none');

            $('#customer-mode').on('select2:select', function(e) {
                var data = e.params.data;

                if (data.id === 'add-new') {
                    $('#new-customer-container').removeClass('d-none');
                    $('#existing-customer-container').addClass('d-none');

                    $('#customer-existing-alert').addClass('d-none');
                    $('#customer-existing-found').empty();
                } else if (data.id === 'search-existing') {
                    $('#existing-customer-container').removeClass('d-none');
                    $('#new-customer-container').addClass('d-none');

                    $('#customer-form').trigger('reset');
                }
                console.log(data);
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
                placeholder: "Pilih Customer",
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
        init: function() {
            this._wizardHandler();
            this._select2Handler();
            this._APIs();
        }
    }
});

