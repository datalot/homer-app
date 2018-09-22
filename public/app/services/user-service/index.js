import angular from 'angular';
import UserService from './user.service';

export default angular.module('hepicApp.services.user', [])
  .factory('UserService', /* @ngInject */ ($http, API) => new UserService($http, API));
