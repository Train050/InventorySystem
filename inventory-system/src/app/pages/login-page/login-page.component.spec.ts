import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginPageComponent } from './login-page.component';

import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';

describe('LoginPageComponent', () => {
  let component: LoginPageComponent;
  let httpTestingController: HttpTestingController;
  let router: Router;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, RouterTestingModule],
      declarations: [LoginPageComponent],
    });

    httpTestingController = TestBed.inject(HttpTestingController);
    router = TestBed.inject(Router);

    const fixture = TestBed.createComponent(LoginPageComponent);
    component = fixture.componentInstance;
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  describe('register()', () => {
    it('should make an HTTP POST request to register the user and redirect to the profile page', () => {
      spyOn(router, 'navigate');

      component.registerFirstName = 'John';
      component.registerLastName = 'Doe';
      component.registerEmail = 'johndoe@example.com';
      component.registerPassword = 'password123';
      component.registerPhoneNumber = '4071231234';

      component.register();

      const req = httpTestingController.expectOne('http://localhost:8080/register');
      expect(req.request.method).toEqual('POST');
      expect(req.request.body).toEqual({
        firstName: 'John',
        lastName: 'Doe',
        email: 'johndoe@example.com',
        password: 'password123',
        phoneNumber: '4071231234'
      });

      req.flush({ jwt: 'mock-jwt-token' });

      expect(localStorage.getItem('token')).toEqual('mock-jwt-token');
      expect(router.navigate).toHaveBeenCalledWith(['profile']);

      expect(component.registerFirstName).toBeNull();
      expect(component.registerLastName).toBeNull();
      expect(component.registerEmail).toBeNull();
      expect(component.registerPassword).toBeNull();
      expect(component.registerPhoneNumber).toBeNull();
    });
  });

  describe('login()', () => {
    it('should make an HTTP POST request to authenticate the user and redirect to the profile page', () => {
      spyOn(router, 'navigate');

      component.loginEmail = 'johndoe@example.com';
      component.loginPassword = 'password123';

      component.login();

      const req = httpTestingController.expectOne('http://localhost:8080/login');
      expect(req.request.method).toEqual('POST');
      expect(req.request.body).toEqual({
        email: 'johndoe@example.com',
        password: 'password123',
      });

      req.flush({ jwt: 'mock-jwt-token' });

      expect(localStorage.getItem('token')).toEqual('mock-jwt-token');
      expect(router.navigate).toHaveBeenCalledWith(['profile']);

      expect(component.loginEmail).toBeNull();
      expect(component.loginPassword).toBeNull();
    });
  });
});
