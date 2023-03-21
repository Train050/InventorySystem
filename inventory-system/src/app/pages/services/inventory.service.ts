import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class InventoryService {

  constructor(private http: HttpClient) { }

  addItems(data: any) {
    const item = {
      ...data
    }
    return this.http.post('api/inventory', item)
  }
}
