import { Component, Inject, OnInit } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";
import { InventoryService } from '../services/inventory.service';

@Component({
    selector: 'app-add-item',
    templateUrl: 'Add-item.html',
    styleUrls: ['add-item.css'],
  })
  export class AddItemComponent implements OnInit{
    form!: FormGroup;
    constructor(
      @Inject(MAT_DIALOG_DATA) public data: {},
      public dialogRef: MatDialogRef<any>,
      private inventoryService: InventoryService
    ) {}

    ngOnInit(): void {
      this.form = new FormGroup({
        productName: new FormControl('', Validators.required),
        dateAcquired: new FormControl('', Validators.required),
        quantity: new FormControl('', Validators.required),
      })
    }

    submitItem() {
      console.log(this.form.value);
      this.inventoryService.addItems(this.form.value).subscribe((res: any) => {
        console.log(res)
        this.dialogRef.close();
      })
      // if (this.form.valid) {
      //   this.dialogRef.close();
      // }
    }
  }
