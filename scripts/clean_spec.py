

def main():
    with open('../shipbob_open_api.json', 'r') as f:
        with open('../shipbob_open_api_fixed.json', 'w') as w:
            for line in f:
                line = replace_if_present(line, '$type', 'ttype')

                line = remove_if_present(line, 'Shipbob.')
                line = remove_if_present(line, 'ShipBob.')
                line = remove_if_present(line, 'Microsoft.AspNetCore.Mvc.')
                line = remove_if_present(line, 'Public.Common.Models.')
                line = remove_if_present(line, 'Public.Common.')
                line = remove_if_present(line, 'Public.Api.Models.')
                line = remove_if_present(line, 'Public.Api.ViewModels.')
                line = remove_if_present(line, 'Api.ViewModels.')
                line = remove_if_present(line, 'ViewModels.')
                line = remove_if_present(line, 'Models.')

                line = remove_if_present(line, 'Presentation.')
                line = remove_if_present(line, 'Integrations.')
                line = remove_if_present(line, 'Location.')
                line = remove_if_present(line, 'Orders.')
                line = remove_if_present(line, 'Channels.')
                line = remove_if_present(line, 'Inventory.')
                line = remove_if_present(line, 'Products.')
                line = remove_if_present(line, 'Receiving.')
                line = remove_if_present(line, 'Returns.')
                line = remove_if_present(line, 'ViewModel')
                line = remove_if_present(line, 'Model')

                line = replace_if_present(line, 'Microsoft_AspNetCore_Mvc_ValidationProblemDetails_allOf', 'ValidationProblemDetails_allOf')

                w.write(line)


def remove_if_present(line, word):
    return replace_if_present(line, word, '')


def replace_if_present(line, word, replace):
    if word in line:
        return line.replace(word, replace)
    return line


if __name__ == "__main__":
    main()
