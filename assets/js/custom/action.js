"use strict";

var Dashboard = {};
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

                APIs.CustomerAddEdit(url, method, payload, function(res) {
                    console.log(res);
                    $('#table-users').DataTable().ajax.reload();
                    $('#userEditModal').modal('hide');
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
                    window.location.replace(base_url+'/dashboard');
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

    Dashboard = {
        _loadServiceTransactionStatus: function() {
            APIs.ServiceTransactionStatusDashboard(function(res) {
                console.log(res);
                $('.status-new').html(res.new);
                $('.status-inprogress').html(res.in_progress);
                $('.status-completed').html(res.completed);
            });
        },
        _loadPieChartDemo: function() {
            // Set new default font family and font color to mimic Bootstrap's default styling
            Chart.defaults.global.defaultFontFamily = 'Nunito', '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
            Chart.defaults.global.defaultFontColor = '#858796';

// Pie Chart Example
            var ctx = document.getElementById("myPieChart");
            var myPieChart = new Chart(ctx, {
                type: 'doughnut',
                data: {
                    labels: ["Direct", "Referral", "Social"],
                    datasets: [{
                        data: [55, 30, 15],
                        backgroundColor: ['#4e73df', '#1cc88a', '#36b9cc'],
                        hoverBackgroundColor: ['#2e59d9', '#17a673', '#2c9faf'],
                        hoverBorderColor: "rgba(234, 236, 244, 1)",
                    }],
                },
                options: {
                    maintainAspectRatio: false,
                    tooltips: {
                        backgroundColor: "rgb(255,255,255)",
                        bodyFontColor: "#858796",
                        borderColor: '#dddfeb',
                        borderWidth: 1,
                        xPadding: 15,
                        yPadding: 15,
                        displayColors: false,
                        caretPadding: 10,
                    },
                    legend: {
                        display: false
                    },
                    cutoutPercentage: 80,
                },
            });
        },
        _loadAreaChartDemo: function() {
            // Set new default font family and font color to mimic Bootstrap's default styling
            Chart.defaults.global.defaultFontFamily = 'Nunito', '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
            Chart.defaults.global.defaultFontColor = '#858796';

            function number_format(number, decimals, dec_point, thousands_sep) {
                // *     example: number_format(1234.56, 2, ',', ' ');
                // *     return: '1 234,56'
                number = (number + '').replace(',', '').replace(' ', '');
                var n = !isFinite(+number) ? 0 : +number,
                    prec = !isFinite(+decimals) ? 0 : Math.abs(decimals),
                    sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
                    dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
                    s = '',
                    toFixedFix = function(n, prec) {
                        var k = Math.pow(10, prec);
                        return '' + Math.round(n * k) / k;
                    };
                // Fix for IE parseFloat(0.55).toFixed(0) = 0;
                s = (prec ? toFixedFix(n, prec) : '' + Math.round(n)).split('.');
                if (s[0].length > 3) {
                    s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);
                }
                if ((s[1] || '').length < prec) {
                    s[1] = s[1] || '';
                    s[1] += new Array(prec - s[1].length + 1).join('0');
                }
                return s.join(dec);
            }

// Area Chart Example
            var ctx = document.getElementById("myAreaChart");
            var myLineChart = new Chart(ctx, {
                type: 'line',
                data: {
                    labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
                    datasets: [{
                        label: "Earnings",
                        lineTension: 0.3,
                        backgroundColor: "rgba(78, 115, 223, 0.05)",
                        borderColor: "rgba(78, 115, 223, 1)",
                        pointRadius: 3,
                        pointBackgroundColor: "rgba(78, 115, 223, 1)",
                        pointBorderColor: "rgba(78, 115, 223, 1)",
                        pointHoverRadius: 3,
                        pointHoverBackgroundColor: "rgba(78, 115, 223, 1)",
                        pointHoverBorderColor: "rgba(78, 115, 223, 1)",
                        pointHitRadius: 10,
                        pointBorderWidth: 2,
                        data: [0, 10000, 5000, 15000, 10000, 20000, 15000, 25000, 20000, 30000, 25000, 40000],
                    }],
                },
                options: {
                    maintainAspectRatio: false,
                    layout: {
                        padding: {
                            left: 10,
                            right: 25,
                            top: 25,
                            bottom: 0
                        }
                    },
                    scales: {
                        xAxes: [{
                            time: {
                                unit: 'date'
                            },
                            gridLines: {
                                display: false,
                                drawBorder: false
                            },
                            ticks: {
                                maxTicksLimit: 7
                            }
                        }],
                        yAxes: [{
                            ticks: {
                                maxTicksLimit: 5,
                                padding: 10,
                                // Include a dollar sign in the ticks
                                callback: function(value, index, values) {
                                    return '$' + number_format(value);
                                }
                            },
                            gridLines: {
                                color: "rgb(234, 236, 244)",
                                zeroLineColor: "rgb(234, 236, 244)",
                                drawBorder: false,
                                borderDash: [2],
                                zeroLineBorderDash: [2]
                            }
                        }],
                    },
                    legend: {
                        display: false
                    },
                    tooltips: {
                        backgroundColor: "rgb(255,255,255)",
                        bodyFontColor: "#858796",
                        titleMarginBottom: 10,
                        titleFontColor: '#6e707e',
                        titleFontSize: 14,
                        borderColor: '#dddfeb',
                        borderWidth: 1,
                        xPadding: 15,
                        yPadding: 15,
                        displayColors: false,
                        intersect: false,
                        mode: 'index',
                        caretPadding: 10,
                        callbacks: {
                            label: function(tooltipItem, chart) {
                                var datasetLabel = chart.datasets[tooltipItem.datasetIndex].label || '';
                                return datasetLabel + ': $' + number_format(tooltipItem.yLabel);
                            }
                        }
                    }
                }
            });
        },
        init: function() {
            this._loadServiceTransactionStatus();
            this._loadPieChartDemo();
            this._loadAreaChartDemo();
        }
    };
});

