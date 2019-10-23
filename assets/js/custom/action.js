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
});

