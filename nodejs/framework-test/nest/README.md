# NestJS: Uma abordagem modernada e estruturada 

NesteJS é uma estrutura conhecida por construir aplicativos nodejs escalonáveis e eficientes no lado do servidor. Ele usa javascirpt progressivo e tem a capacidade de escrever código em typeScirpt. Embora seja totalmente compatível com Typesscript, ele pode escrever código em Javascript puro e incluir programação orientada a objetos, programação funcional e progrmaação reativa funcional. 

## Principais recursos; O que o faz se destarcar 

1. Modularidade 
Nestjs permite que o código seja divido em módulos gerenciáveis separados, o que o torna mais fácil de manter. Por exemplo vejamos o módulo abaixo: 

```
import { Module } from '@nestjs/common';

@Module({
 imports: [
  CacheModule
 ],
 controllers: [PaymentController],
 providers: [PaymentService],
})
export class PaymentModule {}
```

Isso "payment module" pode ser exportado para outros módulos. Neste exemplo, exportamos comum "cache module" dentro deste módulo. Como nest.js possui uma estrutura de módulo, facilita sua manutenção. 

2. Escalável 
Nestejs permite escolanemnteo contínuo, dividino aplicativos em móduloe gerenciáveis, oferecendo suporte à substituição flexível de componentes e acomodando alto tráfego por meio de microsserviços e operações assíncronas. Ele garente o manuseio eficiente do aumento da carga de trabalho, mantendo a confiabilidade. 

3. Injeção de depdência 
A injeção de depedência é simplesmente o método de adicionar uma dependência externa a uma classe, em vez de criá-la dentro da própria classe. vejamos o exemplo: 

```
importar { 
HttpException, Injetável, NotFoundException 
} de '@nestjs/common'; 

@Injectable() 
export class PaymentService { 

construtor() {} 

getReceipt() { 
   return 'Recibo de pagamento'; 
} 

}
```

Criamos "payment service"e adicionamos "@Injectable()" anotaç~eos para torná-lo injetável. Podemos usar o serviço criado conforme indicado abaixo. 

```
importe {Controller, Get, Post, Body} de '@nestjs/common'; 
importar {PaymentService} de './payment.service'; 
@Controller('payment') 
export class PaymentController { 
construtor(private readonly paymentService: PaymentService) {} 
@Get() 
getPaymentReceipt() { 
return this.paymentService.getReceipt(); 
} 
}
```

4. Digite segurança 

Nestjjs usa Typescirpt para fornecer segurança de tipo, que pode ser usada para detectar possível erros durante o desenvolvimento e melhorar a capacidade de manutenção do código. 

```
export class PaymentDto { 

  @IsNotEmpty() 
  @IsEnum(SERVICE_PROVIDER_SLUG, { 
    mensagem: `serviceProvider inválido. As opções válidas são: ${Object.values(SERVICE_PROVIDER_SLUG).join(', ')}`, 
  }) 
  serviceProvider: string; 

  @IsNotEmpty() 
  @IsNumber() 
  valor: número; 

  @IsNotEmpty() 
  @IsString() 
  validadePeriod: string; 

  @IsNotEmpty() 
  @IsArray() 
  @ArrayNotEmpty() 
  @ValidateNested() 
  @Type(() => PaymentAttributesDto) 
  paymentAttributes: PaymentAttributesDto[] 

}
```

Neste exmeplo, criamos um DTO que inclui vários parâmetros e adicionamos anotações para validar o tipo de parâmetro. Por exemplo, se enviarmos um valor de string para o parâmetro "valor"  ocorrerá um erro. 


