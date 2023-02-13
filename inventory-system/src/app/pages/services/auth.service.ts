import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor() { }

  login(uname: string, pword: string) {
    if (uname === 'username' && pword === '1234') {
      return 200;
    } else {
      return 403;
    }
  }

  registration(uname: string, pword: string, email: string, phoneNumber: string)
  {
    if (uname === '' || pword === '' || email === '' || phoneNumber === '')
    {
      return 403;
    }
    else
    {
      return 200;
    }
  }
}
