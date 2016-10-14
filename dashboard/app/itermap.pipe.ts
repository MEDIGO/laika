import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'itermap'
})
export class ItermapPipe {
  transform(value: Object): Array<Object> {
    var res = [];

    for (var key in value) {
      if (value.hasOwnProperty(key)) {
        res.push({key: key, value: value[key]});
      }
    }

    return res;
  }
}
