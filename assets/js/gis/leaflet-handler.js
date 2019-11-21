"use strict";

var LeafletMap = {};

$(function() {
    LeafletMap = {
        _basicMap: function() {
            var mymap = L.map('basic-map').setView([51.505, -0.09], 13);
            L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token=pk.eyJ1IjoiZmFkbGlrYWRuIiwiYSI6ImNrMzgzdm0xYTA0MHkzYm1ncm9xampnMDEifQ.p5iIhIFCDOeqKy5sVomI2Q', {
                maxZoom: 18,
                // attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, ' +
                //     '<a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
                //     'Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
                id: 'mapbox.streets',
                accessToken: 'pk.eyJ1IjoiZmFkbGlrYWRuIiwiYSI6ImNrMzgzdm0xYTA0MHkzYm1ncm9xampnMDEifQ.p5iIhIFCDOeqKy5sVomI2Q',
            }).addTo(mymap);

            $(window).on('resize', function() {
                $('#basic-map').height($(window).height() - 150);
                mymap.invalidateSize()
            }).trigger('resize');

        },
        init: function() {
            this._basicMap();
        }
    };
});