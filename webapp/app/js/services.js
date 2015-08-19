'use strict';
/* global localStorage */

/* Services */


angular.module('myApp.services',[])
    .factory('lstorage', function() {
        
        var storageService = {};
        storageService.name = 'mswift42themecreator';
        storageService.containsKey = function() {
            return localStorage.getItem(storageService.name) !== null;
        };
        storageService.loadFaces = function() {
            return JSON.parse(localStorage.getItem(storageService.name));
        };
        storageService.setFaces = function(faces) {
            localStorage.setItem(storageService.name, JSON.stringify(faces));
        };
        storageService.deleteKey = function() {
            localStorage.removeItem(storageService.name);
        };
        return storageService;
    });
angular.module('myApp.services')
    .factory('presets',function() {
    // .factory('presets', function() {

        var presetService = {};
        presetService.warmnight = function () {
            return {deffacefg : "#b1b1b1", deffacebg : "#292424",
                        keywordface : "#96905f", builtinface : "#71a46c",
                        stringface : "#71a19f", functionnameface : "#b680b1",
                        typeface : "#8b8fc6", constantface : "#bd845f",
                        variableface : "#c27d7b", warningface : "#e81050",
                    commentface : "#5d5a58"
                   };
        };
        presetService.oldlace = function() {
            return {deffacefg : "#525252", deffacebg : "#fdf5e6",
                    keywordface : "#3f567b", builtinface : "#7b4135",
                        stringface : "#305f5e", functionnameface : "#714355",
                        typeface : "#634575", constantface : "#64502f",
                        variableface : "#3f5b32", warningface : "#fa0c0c",
                    commentface : "#949494"
                   };
        };
        presetService.lightsoap = function() {
            return {
                deffacefg : "#474747", deffacebg : "#fafad4",
                keywordface : "#4f4b78", builtinface : "#365f71",
                        stringface : "#355d3e", functionnameface : "#86546c",
                        typeface : "#5a532d", constantface : "#794732",
                        variableface : "#5f3041", warningface : "#f70f2a",
                commentface : "#a1a1a1"
            };
        };
        presetService.greymatters = function() {
            return {
                                deffacefg : "#2f2f2f", deffacebg : "#f9fbfd",
                keywordface : "#3f567b", builtinface : "#7b4135",
                        stringface : "#305f5e", functionnameface : "#714355",
                        typeface : "#634575", constantface : "#64502f",
                        variableface : "#3f5b32", warningface : "#fa0c0c",
                commentface : "#949494"
            };
        };
        presetService.softcharcoal = function() {
            return {
                                deffacefg : "#c2c2c2", deffacebg : "#191919",
                keywordface : "#8aa234", builtinface : "#54686d",
                        stringface : "#5d90cd", functionnameface : "#7a8bbd",
                        typeface : "#8aa6c1", constantface : "#8aa6c1",
                        variableface : "#8885b2", warningface : "#ff6523",
                commentface : "#808080"
            };
        };
        presetService.softmorning = function() {
                        return {
                                deffacefg : "#282828", deffacebg : "#f2f1f0",
                keywordface : "#8aa234", builtinface : "#727170",
                        stringface : "#3450a2", functionnameface : "#a82e4d",
                        typeface : "#727170", constantface : "#f03f3f",
                        variableface : "#f03f3f", warningface : "#ff6523",
                commentface : "#808080"
                        };
        };
        presetService.softstone = function() {
                          return {
                                deffacefg : "#000000", deffacebg : "#efece3",
                keywordface : "#490026", builtinface : "#9e0045",
                        stringface : "#0f126e", functionnameface : "#340557",
                        typeface : "#014500", constantface : "#374014",
                        variableface : "#014500", warningface : "#ff6523",
                commentface : "#373737"
                          };
        };
                
        return presetService;
    
    });
