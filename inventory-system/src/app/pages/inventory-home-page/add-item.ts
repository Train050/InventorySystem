import { Component, Inject, OnInit } from "@angular/core";
import { FormControl, FormGroup, Validators } from "@angular/forms";
import { MatDialogRef, MAT_DIALOG_DATA } from "@angular/material/dialog";

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
    ) {}

    ngOnInit(): void {
      this.form = new FormGroup({
        id: new FormControl('', Validators.required),
        productName: new FormControl('', Validators.required),
        dateAcquired: new FormControl('', Validators.required),
        quantity: new FormControl('', Validators.required),
      })
    }

    submitItem() {
      console.log(this.form.value);
      if (this.form.valid) {
        this.dialogRef.close();
      }
    }
  }
  