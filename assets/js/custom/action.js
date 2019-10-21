"use strict";

const base_url = 'http://localhost:8080';

var Login = {};
var Register = {};
var Users = {};

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
                        "data": "id", render: function(data) {
                            return `<a href='#' class='btn btn-sm btn-success user_edit' id-key=${data} >Edit</a> &nbsp; <a href='#' class='btn btn-sm btn-danger user_delete' id-key=${data}>Delete</a>`;
                        }
                    }
                ]
            });
        },
        _APIs: function() {

        },
        _loadUsersAPI: function(callback) {
            // var self = this;
            // $.ajax({
            //     url: base_url + "/api/session/users",
            //     method: 'GET',
            //     dataType: 'json',
            //     contentType: 'application/json',
            //     success: callback,
            // }).catch(function(err) {
            //     console.log(err);
            // });
        },
        _loadPage: function() {
            this._loadUsersAPI(function(res) {
                // console.log(res);
                // var $userTable = $('.users-data-container');
                // $userTable.empty();
                // var usersData = "";
                // res.forEach(function(item, index) {
                //     var userData = `<tr>
                //                         <td>${item.name}</td>
                //                         <td>${item.email}</td>
                //                         <td>${item.phone}</td>
                //                         <td>${item.company}</td>
                //                         <td>${item.IsActive}</td>
                //                     </tr>`;
                //     usersData = usersData + userData;
                // });
                // $userTable.html(usersData);
            });
        },
        init: function() {
            this._datatables();
            this._APIs();
            this._loadPage();
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
});

