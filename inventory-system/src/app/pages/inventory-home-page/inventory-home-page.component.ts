import { Component } from '@angular/core';

export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
}

const ELEMENT_DATA: PeriodicElement[] = [
  {position: 1, name: 'Hydrogen', weight: 1.0079, symbol: '13250'},
  {position: 2, name: 'Helium', weight: 4.0026, symbol: '4215'},
  {position: 3, name: 'Lithium', weight: 6.941, symbol: '5871'},
  {position: 4, name: 'Beryllium', weight: 9.0122, symbol: '1688'},
  {position: 5, name: 'Boron', weight: 10.811, symbol: '2584'},
  {position: 6, name: 'Carbon', weight: 12.0107, symbol: '354'},
  {position: 7, name: 'Nitrogen', weight: 14.0067, symbol: '2351'},
  {position: 8, name: 'Oxygen', weight: 15.9994, symbol: '1684'},
  {position: 9, name: 'Fluorine', weight: 18.9984, symbol: '18548'},
  {position: 10, name: 'Neon', weight: 20.1797, symbol: '8135'},
];

@Component({
  selector: 'app-inventory-home-page',
  templateUrl: './inventory-home-page.component.html',
  styleUrls: ['./inventory-home-page.component.scss']
})
export class InventoryHomePageComponent {
  displayedColumns: string[] = ['demo-position', 'demo-name', 'demo-weight', 'demo-symbol'];
  dataSource = ELEMENT_DATA;
}
